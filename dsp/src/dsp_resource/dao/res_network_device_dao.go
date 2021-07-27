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
//查询网络设备列表
func ResNetworkDeviceList(pojo cusfun.ParamsPOJO) (res []model.ResNetworkDevicePage, err error, total int) {
	var count int64
	db := util.DbConn.Table("(select  res.id,res.res_name,res.res_ip_v4,res.res_ip_v6,res.dic_connect_id,res.dic_type_id,res.dic_city_id,res.dic_manufactor_id,res.department_id,res.res_manager_user_id,res.`status`,res.z_desc,date_format(res.create_date, '%Y-%m-%d %H:%i:%s') as create_date,res.create_user_id,date_format(res.modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,res.modify_user_id,res.res_password_policy_id,res.secure_port,res.somf_ip,sec.res_id,sec.network_device_type,sec.admin_account,sec.admin_password,sec.root_password,sec.admin_con_sym,sec.root_con_sym,sec.is_return,sec.is_kvm,sec.res_skip_id,sec.protocol_type,sec.protocol_port,sec.dic_version_id,sec.dic_model_id,deviceType.dic_name as network_device_type_name,department.dep_name dep_name,policy.policy_name policy_name,skip.res_name skip_name,USER.uname  from z_res res left join z_res_network_device sec on sec.res_id=res.id left join z_dictionary deviceType on deviceType.id=sec.network_device_type LEFT JOIN z_department department ON department.id = res.department_id LEFT JOIN z_password_policy policy ON policy.id = res.res_password_policy_id LEFT JOIN z_res skip ON skip.id = sec.res_skip_id LEFT JOIN z_user USER ON USER.uid = res.res_manager_user_id) t").Select(" id,res_name,res_ip_v4,res_ip_v6,dic_connect_id,dic_type_id,dic_city_id,dic_manufactor_id,department_id,res_manager_user_id,`status`,z_desc,create_date,create_user_id,modify_date,modify_user_id,res_password_policy_id,secure_port,somf_ip,res_id,network_device_type,admin_account,admin_password,root_password,admin_con_sym,root_con_sym,is_return,is_kvm,res_skip_id,protocol_type,protocol_port,dic_version_id,dic_model_id,network_device_type_name,dep_name,policy_name,skip_name,uname")
	db = db.Where("dic_type_id = ?", 1303)
	db = db.Where("status != 0")
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}

//添加网络设备
func CreateResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	var list = []string{"SSH", "TELNET"}
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	resInfo.Res.DicTypeId = 1303
	resInfo.Res.Status = 1
	errMain := tx.Create(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.ResNetworkDevice.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Create(&resInfo.ResNetworkDevice).Error

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
		}
		errpro = tx.Create(&protocol).Error
	}
	if errMain == nil && err == nil && err == errpro {
		count = 1
		tx.Commit()
		fmt.Println("添加事务成功")
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("添加事务回退")
	}
	return resInfo, int(count), err
}

//修改网络设备
func UpdateResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	errMain := tx.Save(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.ResNetworkDevice.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Save(&resInfo.ResNetworkDevice).Error

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

//删除网络设备
func DeleteResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	count := util.DbConn.Table("z_res").Where("id=?", resInfo.Res.Id).Updates(map[string]interface{}{"status": 0, "department_id": 3}).RowsAffected
	return resInfo, int(count), nil
}

//批量添加数据库载体（导入）
func CreateNetworkMore(url string) ([]model.HostLogs, error, int) {
	var ips = ""
	var names = ""
	var number = 0
	var str = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	var a []model.ResNetworkDeviceInfo
	var b []model.Res
	util.DbConn.Table("z_res left JOIN z_res_store ON z_res.id = z_res_store.res_id").Select("z_res.*,z_res_store.*").Find(&b).Order("Id desc").Where("z_res.status != 0")
	util.DbConn.Table("z_res").Select("z_res.*").Find(&a).Order("Id desc")
	fmt.Println("查询结果a  ", a)
	fmt.Println("查询结果b  ", b)
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
	var Network = make([]model.ResNetworkDevice, 0)
	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
		return hostLogs, err, 0
	}
	var resSmall = new(model.Res)
	var NetworkSmall = new(model.ResNetworkDevice)
	rows, err := f.GetRows("网络设备模板")
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
				logs.ErrorLog += "网络设备类型不能为空 "
			}
			if rows[i][9] == "" {
				can = false
				logs.ErrorLog += "是否回车不能为空 "
			}
			if rows[i][10] == "" {
				can = false
				logs.ErrorLog += "是否kvm不能为空 "
			}
			if rows[i][11] == "" {
				can = false
				logs.ErrorLog += "协议类型不能为空 "
			}
			if rows[i][12] == "" {
				can = false
				logs.ErrorLog += "协议端口不能为空 "
			}
			if rows[i][13] == "" {
				can = false
				logs.ErrorLog += "城市不能为空 "
			}
		}
		if can {
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
			if rows[i][9] == "是" {
				NetworkSmall.IsReturn = 1
			} else {
				NetworkSmall.IsReturn = 0
			}
			if rows[i][10] == "是" {
				NetworkSmall.IsKvm = 1
			} else {
				NetworkSmall.IsKvm = 0
			}
			NetworkSmall.ResId = number
			NetworkSmall.ProtocolType, _ = strconv.Atoi(Substr(rows[i][11], strings.Index(rows[i][11], "_")))
			NetworkSmall.ProtocolPort, _ = strconv.Atoi(rows[i][12])
			NetworkSmall.NetworkDeviceType, _ = strconv.Atoi(Substr(rows[i][8], strings.Index(rows[i][8], "_")))
			resSmall.DicManufactorId, _ = strconv.Atoi(Substr(rows[i][7], strings.Index(rows[i][7], "_")))
			resSmall.DicCityId, _ = strconv.Atoi(Substr(rows[i][13], strings.Index(rows[i][13], "_")))
			resSmall.Id = number
			resSmall.DicTypeId = 1303
			Network = append(Network, *NetworkSmall)
			res = append(res, *resSmall)
			logs.ErrorLog += "无异常 "
		}
		hostLogs = append(hostLogs, *logs)
	}
	countMain := util.DbConn.Table("z_res").Create(&res).RowsAffected
	count := util.DbConn.Table("z_res_network_device").Create(&Network).RowsAffected
	fmt.Println(count)
	return hostLogs, err, int(countMain)
}

/*
导出网络设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportDevice(resInfo model.ResNetworkDeviceInfo) ([]model.ExportNetwork, error, int) {
	var count int64
	var exportList []model.ExportNetwork
	db := util.DbConn.Table("z_res res left join z_res_network_device sec on sec.res_id=res.id  left join z_dictionary dicType on dicType.id=res.dic_type_id left join z_dictionary manufactor on manufactor.id=res.dic_manufactor_id  left join z_dictionary city on city.id=res.dic_city_id  left join z_department department on department.id=res.department_id left join z_res_account account on account.res_id=res.id  left join z_password_policy policy on policy.id=res.res_password_policy_id  left join z_dictionary deviceType on deviceType.id=sec.network_device_type left join z_res skip on skip.id=sec.res_skip_id left join z_dictionary protocolType on protocolType.id=sec.protocol_type").Select("res.res_name,res.res_ip_v4,res.res_ip_v6,dicType.dic_name as res_type_name,department.dep_name as dep_name,account.res_account,account.res_account_name,policy.policy_name as res_password_policy,if(res.dic_connect_id=1,'是','否') as dic_connect_id,manufactor.dic_name as manufactor_name,deviceType.dic_name as network_device_type,if(sec.is_return=1,'是','否') as is_return,if(sec.is_kvm=1,'是','否') as is_kvm,skip.res_name as res_skip_name,protocolType.dic_name as protocol_type,sec.protocol_port,city.dic_name as dic_city")
	db = db.Where(" res.dic_type_id = ?", 1303)
	db = db.Where(" res.status != ?", 0)
	if resInfo.Res.ResIpV4 != "" {
		db = db.Where("res.res_ip_v4 = ?", resInfo.Res.ResIpV4)
	}
	if resInfo.Res.ResName != "" {
		db = db.Where(" res.res_name like (?)", "%"+resInfo.Res.ResName+"%")
	}
	if resInfo.ResNetworkDevice.NetworkDeviceType != 0 {
		db = db.Where(" sec.network_device_type = ?", resInfo.ResNetworkDevice.NetworkDeviceType)
	}
	if resInfo.Res.DepartmentId != 0 {
		db = db.Where(" res.department_id = ?", resInfo.Res.DepartmentId)
	}
	db.Count(&count).Order("res.id desc").Find(&exportList)
	return exportList, nil, int(count)
}
