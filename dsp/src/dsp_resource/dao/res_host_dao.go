package dao

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"regexp"

	"strconv"
	"strings"
)

/**
 * @Author: wangyue
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/3/23
 */
func HostList(pojo cusfun.ParamsPOJO) (res []model.HostPage, err error, total int) {
	var count int64
	db := util.DbConn.Table("(select res.id,res.res_name,res.res_ip_v4,res.res_ip_v6,res.dic_connect_id,res.dic_type_id,res.dic_city_id,res.dic_manufactor_id,res.department_id,res.res_manager_user_id,res.`status`,res.z_desc,date_format(res.create_date, '%Y-%m-%d %H:%i:%s') as create_date,res.create_user_id,date_format(res.modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,res.modify_user_id,res.res_password_policy_id,res.secure_port,res.somf_ip,sec.res_id,sec.host_type,sec.host_system_type,sec.host_system_version,sec.protocol_type,sec.protocol_port,sec.is_sudo,sec.admin_account,sec.admin_password,sec.admin_con_sym,sec.unix_su_type,sec.unix_root_password,sec.unix_root_con_sym,sec.unix_res_skip_id,sec.win_type,sec.win_domain_name,sec.win_domain_dn,sec.win_domain_res_id,sec.dic_version_id,sec.dic_model_id,sec.rdp_port,hostType.dic_name as host_type_name,department.dep_name dep_name ,policy.policy_name  policy_name,skip.res_name skip_name,user.uname,domain.res_name as domain_name from z_res res left JOIN z_res_host sec ON res.id = sec.res_id left join z_dictionary hostType on hostType.id=sec.host_type left join z_department department on department.id=res.department_id  LEFT join z_password_policy policy on policy.id=res.res_password_policy_id LEFT JOIN z_res skip ON skip.id = sec.unix_res_skip_id LEFT join z_user user on user.uid=res.res_manager_user_id left join z_res domain on domain.id=sec.win_domain_res_id) t").Select("id,res_name,res_ip_v4,res_ip_v6,dic_connect_id,dic_type_id,dic_city_id,dic_manufactor_id,department_id,res_manager_user_id,`status`,z_desc,create_date,create_user_id,modify_date,modify_user_id,res_password_policy_id,secure_port,somf_ip,res_id,host_type,host_system_type,host_system_version,protocol_type,protocol_port,is_sudo,admin_account,admin_password,admin_con_sym,unix_su_type,unix_root_password,unix_root_con_sym,unix_res_skip_id,win_type,win_domain_name,win_domain_dn,win_domain_res_id,dic_version_id,dic_model_id,rdp_port,host_type_name,dep_name,policy_name,skip_name,uname,domain_name")
	db = db.Where("dic_type_id = ?", 1301)
	db = db.Where("status != 0")
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}

//添加主机载体
func CreateHost(resInfo model.HostInfo) (model.HostInfo, error, int) {
	var list []string
	if resInfo.Host.HostType == 130101 {
		list = []string{"SSH", "SFTP"}
	} else {
		list = []string{"RDP"}
	}
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	resInfo.Res.DicTypeId = 1301
	resInfo.Res.Status = 1
	errMain := tx.Create(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.Host.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Create(&resInfo.Host).Error
	//向协议表里添加数据
	var errpro error
	for _, value := range list {
		var protocol model.HostProtocol
		protocol.ResId = resInfo.Res.Id
		if value == "SSH" {
			protocol.ProtocolPort = 22
			protocol.ProtocolType = 2001
		} else if value == "SFTP" {
			protocol.ProtocolPort = 22
			protocol.ProtocolType = 2003
		} else if value == "TELNET" {
			protocol.ProtocolPort = 23
			protocol.ProtocolType = 2002
		} else if value == "RDP" {
			protocol.ProtocolPort = 3389
			protocol.ProtocolType = 2005
		}
		errpro = tx.Create(&protocol).Error
	}

	if errMain == nil && err == nil && errpro == nil {
		count = 1
		tx.Commit()
		fmt.Println("添加事务成功")
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("添加事务回退")
	}
	return resInfo, err, int(count)
}

//修改主机载体
func UpdateHost(resInfo model.HostInfo) (model.HostInfo, int, error) {
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	errMain := tx.Save(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.Host.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Save(&resInfo.Host).Error
	if errMain == nil && err == nil {
		count = 1
		tx.Commit()
		fmt.Println("修改事务成功")
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("修改事务回退")
	}
	return resInfo, int(count), err
}

//删除主机载体
func DeleteHost(resInfo model.HostInfo) (model.HostInfo, int, error) {
	//修改状态为0
	count := util.DbConn.Table("z_res").Where("id=?", resInfo.Res.Id).Updates(map[string]interface{}{"status": 0, "department_id": 3}).RowsAffected
	return resInfo, int(count), nil
}

//批量添加主机载体（导入）
func CreateHostMore(url string) ([]model.HostLogs, error, int) {
	var ips = ""
	var names = ""
	var number = 0
	var str = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	//全数据查询
	var a []model.HostAll
	//未删除查询，，为了要最后一个id
	var b []model.Res
	util.DbConn.Table("z_res left JOIN z_res_host ON z_res.id = z_res_host.res_id").Select("z_res.*,z_res_host.*").Find(&b).Order("Id desc").Where("z_res.status != 0")
	util.DbConn.Table("z_res").Select("z_res.*").Find(&a).Order("Id desc")
	if len(b) > 0 {
		number = b[len(b)-1].Id
	}
	if len(a) > 0 {
		for i := 0; i < len(a); i++ {
			ips = ips + a[i].ResIpV4 + "@"
			names = names + a[i].ResName + "@"
		}
	}
	var hostLogs = make([]model.HostLogs, 0)
	var res = make([]model.Res, 0)
	var host = make([]model.Host, 0)
	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
		return hostLogs, err, 0
	}

	var resSmall = new(model.Res)
	var hostSmall = new(model.Host)

	rows, err := f.GetRows("主机信息模板")
	//从1开始  第一行不要,for结束以后rows就是校验完的
	for i := 1; i < len(rows); i++ {
		rows[i] = append(rows[i], str...)
		var logs = new(model.HostLogs)
		logs.ResName = rows[i][0]
		logs.ResIpV4 = rows[i][1]
		logs.ErrorLog = string(i) + ":"
		can := true
		fmt.Println("names  ", names)
		//必有字段有无校验
		if len(rows[i]) == 0 {
			can = false
			logs.ErrorLog += "空行 "
		} else {
			if rows[i][0] == "" {
				can = false
				logs.ErrorLog += "资产名称不能为空 "
			} else {
				//资产名重复校验
				if strings.Index(names, rows[i][0]) > -1 {
					can = false
					logs.ErrorLog += "资产名称重复 "
				} else {
					names = names + rows[i][0] + "@"
				}
			}
			if rows[i][1] == "" {
				can = false
				logs.ErrorLog += "Ip不能为空 "
			} else {
				//ip重复校验，后续补充正则，补充ipv6正则
				if strings.Index(ips, rows[i][1]) > -1 {
					can = false
					logs.ErrorLog += "Ip重复 "
				} else {
					regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
					match, _ := regexp.MatchString(regStr, rows[i][1])
					if !match {
						can = false
						logs.ErrorLog += "资产IP格式错误 "
					} else {
						ips = ips + rows[i][1] + "@"
					}
				}
			}
			if rows[i][3] == "" {
				can = false
				logs.ErrorLog += "部门名称不能为空 "
			}
			if rows[i][4] == "" {
				can = false
				logs.ErrorLog += "管理用户名称不能为空 "
			}
			if rows[i][5] == "" {
				can = false
				logs.ErrorLog += "密码策略不能为空 "
			}
			if rows[i][6] == "" {
				can = false
				logs.ErrorLog += "是否可连接不能为空 "
			}
			if rows[i][7] == "" {
				can = false
				logs.ErrorLog += "生产厂家不能为空 "
			}
			if rows[i][8] == "" {
				can = false
				logs.ErrorLog += "主机类型不能为空 "
			}
			if rows[i][9] == "" {
				can = false
				logs.ErrorLog += "win主机类型不能为空 "
			}
			if rows[i][10] == "" {
				can = false
				logs.ErrorLog += "主机系统类型不能为空 "
			}
			if rows[i][11] == "" {
				can = false
				logs.ErrorLog += "主机系统版本不能为空 "
			}
			if rows[i][12] == "" {
				can = false
				logs.ErrorLog += "协议类型不能为空 "
			}
			if rows[i][13] == "" {
				can = false
				logs.ErrorLog += "协议端口不能为空 "
			}
			if rows[i][14] == "" {
				can = false
				logs.ErrorLog += "域名不能为空 "
			}
			if rows[i][15] == "" {
				can = false
				logs.ErrorLog += "域控主机dn名称不能为空 "
			}
			if rows[i][16] == "" {
				can = false
				logs.ErrorLog += "域主机名称不能为空 "
			}
			if rows[i][17] == "" {
				can = false
				logs.ErrorLog += "城市不能为空 "
			}
		}
		if can {
			fmt.Println("i  ", i)
			number++
			resSmall.ResManagerUserId = Substr(rows[i][4], strings.Index(rows[i][4], "_"))
			resSmall.DepartmentId, _ = strconv.Atoi(Substr(rows[i][3], strings.Index(rows[i][3], "_")))
			resSmall.ResPasswordPolicyId, _ = strconv.Atoi(Substr(rows[i][5], strings.Index(rows[i][5], "_")))
			resSmall.ResName = rows[i][0]
			resSmall.ResIpV4 = rows[i][1]
			resSmall.ResIpV6 = rows[i][2]
			resSmall.Status = 1
			if rows[i][6] == "是" {
				resSmall.DicConnectId = 1
			} else {
				resSmall.DicConnectId = 0
			}
			resSmall.DicManufactorId, _ = strconv.Atoi(Substr(rows[i][7], strings.Index(rows[i][7], "_")))
			hostSmall.HostType, _ = strconv.Atoi(Substr(rows[i][8], strings.Index(rows[i][8], "_")))
			hostSmall.WinType, _ = strconv.Atoi(Substr(rows[i][9], strings.Index(rows[i][9], "_")))
			hostSmall.HostSystemType, _ = strconv.Atoi(Substr(rows[i][10], strings.Index(rows[i][10], "_")))
			hostSmall.HostSystemVersion, _ = strconv.Atoi(Substr(rows[i][11], strings.Index(rows[i][11], "_")))
			hostSmall.ProtocolType, _ = strconv.Atoi(Substr(rows[i][12], strings.Index(rows[i][12], "_")))
			hostSmall.ProtocolPort, _ = strconv.Atoi(rows[i][13])
			hostSmall.WinDomainName = rows[i][14]
			hostSmall.WinDomainDn = rows[i][15]
			hostSmall.WinDomainResId, _ = strconv.Atoi(Substr(rows[i][16], strings.Index(rows[i][16], "_")))
			resSmall.DicCityId, _ = strconv.Atoi(Substr(rows[i][17], strings.Index(rows[i][17], "_")))
			resSmall.Id = number
			hostSmall.ResId = number
			hostSmall.UnixResSkipId = number
			hostSmall.WinDomainResId = number
			resSmall.DicTypeId = 1301
			host = append(host, *hostSmall)
			res = append(res, *resSmall)
			logs.ErrorLog += "无异常 "
		}
		hostLogs = append(hostLogs, *logs)
	}
	//保存主表信息
	countMain := util.DbConn.Table("z_res").Create(&res).RowsAffected
	//保存子表信息
	count := util.DbConn.Table("z_res_host").Create(&host).RowsAffected
	fmt.Println(count)
	return hostLogs, err, int(countMain)
}

/*
导出主机设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportHost(resInfo model.HostInfo) ([]model.ExportHost, error, int) {
	var count int64
	var exportList []model.ExportHost
	db := util.DbConn.Table("z_res res left join z_res_host sec on sec.res_id=res.id left join z_dictionary dicType on dicType.id=res.dic_type_id left join z_dictionary manufactor on manufactor.id=res.dic_manufactor_id  left join z_dictionary city on city.id=res.dic_city_id  left join z_department department on department.id=res.department_id left join z_res_account account on account.res_id=res.id  left join z_password_policy policy on policy.id=res.res_password_policy_id  left join z_dictionary hostType on hostType.id=sec.host_type left join z_dictionary winType on winType.id=sec.win_type left join z_dictionary systemType  on systemType.id=sec.host_system_type left join z_dictionary protocolType on protocolType.id=sec.protocol_type left join z_res skip on skip.id=sec.unix_res_skip_id left join z_res domain on domain.id=sec.win_domain_res_id left join z_dictionary version on version.id=sec.dic_version_id left join z_dictionary model on model.id=sec.dic_model_id left join z_dictionary systemVersion on systemVersion.id=sec.host_system_version").Select("res.res_name,res.res_ip_v4,res.res_ip_v6,dicType.dic_name as res_type_name,department.dep_name as dep_name,account.res_account,account.res_account_name,policy.policy_name as res_password_policy,if(res.dic_connect_id=1,'是','否') as dic_connect_id,manufactor.dic_name as manufactor_name,hostType.dic_name as host_type,winType.dic_name as win_type,systemType.dic_name as host_system_type,systemVersion.dic_name host_system_version,protocolType.dic_name as protocol_type,sec.protocol_port as protocol_port,skip.res_name as unix_res_skip_name,sec.win_domain_name,sec.win_domain_dn,domain.res_name as win_domain_res_name,version.dic_name as dic_version_id,model.dic_name as dic_model_id,city.dic_name as dic_city")
	db = db.Where(" res.dic_type_id = ?", 1301)
	db = db.Where(" res.status != ?", 0)
	if resInfo.Res.ResIpV4 != "" {
		db = db.Where("res.res_ip_v4 = ?", resInfo.Res.ResIpV4)
	}
	if resInfo.Res.ResName != "" {
		db = db.Where(" res.res_name like (?)", "%"+resInfo.Res.ResName+"%")
	}
	if resInfo.Host.HostType != 0 {
		db = db.Where(" sec.host_type = ?", resInfo.Host.HostType)
	}
	if resInfo.Res.DepartmentId != 0 {
		db = db.Where(" res.department_id = ?", resInfo.Res.DepartmentId)
	}
	db.Count(&count).Order("res.id desc").Find(&exportList)
	return exportList, nil, int(count)
}

//测试用户名和密码是否可以登录
func LoginHost(params map[string]interface{}) (err error) {
	var cli = new(util.Cli)
	cli.User = params["resAccount"].(string)
	cli.Pwd = params["resAccountPassword"].(string)
	cli.Addr = params["resIpV4"].(string) + ":22"

	var out string
	out, err = cli.Run("pwd")
	fmt.Println(out)
	return
}
