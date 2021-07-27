package dao

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"errors"
	"fmt"
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

func StoreList(pojo cusfun.ParamsPOJO) (res []model.StorePage, err error, total int) {
	var count int64
	db := util.DbConn.Table("(select res.id,res.res_name,IFNULL(host.res_ip_v4,res.res_ip_v4) as res_ip_v4,res.res_ip_v6,res.dic_connect_id,res.dic_type_id,res.dic_city_id,res.dic_manufactor_id,res.department_id,res.res_manager_user_id,res.`status`,res.z_desc,date_format(res.create_date, '%Y-%m-%d %H:%i:%s') as create_date,res.create_user_id,date_format(res.modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,res.modify_user_id,res.res_password_policy_id,res.secure_port,res.somf_ip,sec.res_id,sec.store_name,sec.store_type,sec.store_version,sec.store_port,sec.res_host_id,sec.store_class,sec.store_url,sec.store_sid,sec.store_admin_account,sec.store_admin_account_password,storeType.dic_name as store_type_name,department.dep_name  dep_name,policy.policy_name policy_name,host.res_name host_name,USER.uname from z_res res left JOIN z_res_store sec ON res.id = sec.res_id left join z_dictionary storeType on storeType.id=sec.store_type LEFT JOIN z_department department ON department.id = res.department_id LEFT JOIN z_password_policy policy ON policy.id = res.res_password_policy_id LEFT JOIN z_res host ON host.id = sec.res_host_id LEFT JOIN z_user USER ON USER.uid = res.res_manager_user_id ) t").Select(" id,res_name,res_ip_v4,res_ip_v6,dic_connect_id,dic_type_id,dic_city_id,dic_manufactor_id,department_id,res_manager_user_id,`status`,z_desc,create_date,create_user_id,modify_date,modify_user_id,res_password_policy_id,secure_port,somf_ip,res_id,store_name,store_type,store_version,store_port,res_host_id,store_class,store_url,store_sid,store_admin_account,store_admin_account_password,store_type_name,dep_name,policy_name,host_name,uname ")
	db = db.Where("dic_type_id = ?", 1302)
	db = db.Where("status != 0")
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}

//添加主机载体
func CreateStore(resInfo model.StoreInfo) (model.StoreInfo, error, int) {
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//获取数据库中的最大安全端口
	maxport := 9999
	util.DbConn.Table("z_res").Select("max(secure_port)").Take(&maxport)
	if maxport == 0 {
		maxport = 10000
	} else {
		maxport = maxport + 1
	}

	fmt.Println(maxport)
	//保存主表信息
	resInfo.Res.DicTypeId = 1302
	resInfo.Res.Status = 1
	resInfo.Res.SecurePort = maxport
	errMain := tx.Create(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.Store.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Create(&resInfo.Store).Error
	//向协议表中添加一条数据
	var protocol model.HostProtocol
	protocol.ResId = resInfo.Res.Id
	protocol.ProtocolPort = resInfo.Store.StorePort
	protocol.ProtocolType = 2006
	errpro := tx.Create(&protocol).Error

	AddIpRules(util.Cfg.IptPath.IptRules, resInfo, maxport)
	if errMain == nil && err == nil && errpro == nil {
		count = 1
		tx.Commit()
		fmt.Println("添加事务成功")
		updateIptRulesCmd := "/opt/somf/dgae/bin/reloadipt"
		_, err = util.ExecShell(updateIptRulesCmd)
		if err != nil {
			err = errors.New("数据库访问策略重载失败")
		}
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("添加事务回退")
	}
	return resInfo, err, int(count)
}

func UpdateStore(resInfo model.StoreInfo) (model.StoreInfo, int, error) {
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	errMain := tx.Save(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.Store.ResId = int(resInfo.Res.Id)
	//保存子表信息
	err := tx.Save(&resInfo.Store).Error

	errpro := util.DbConn.Table("z_res_host_protocol").Where("res_id=? and protocol_type=2006 ", resInfo.Res.Id).Updates(map[string]interface{}{"protocol_port": resInfo.Store.StorePort}).Error
	UpdateIpRules(util.Cfg.IptPath.IptRules, resInfo, resInfo.Res.SecurePort)
	if errMain == nil && err == nil && errpro == nil {
		count = 1
		tx.Commit()
		fmt.Println("修改事务成功")
		updateIptRulesCmd := "/opt/somf/dgae/bin/reloadipt"
		_, err = util.ExecShell(updateIptRulesCmd)
		if err != nil {
			err = errors.New("ipt_rules文件修改失败")
		}
	} else {
		count = 0
		tx.Rollback()
		fmt.Println("修改事务回退")
	}
	return resInfo, int(count), err
}

//删除主机载体
func DeleteStore(resInfo model.StoreInfo) (model.StoreInfo, int, error) {
	//修改状态为0
	count := util.DbConn.Table("z_res").Where("id=?", resInfo.Res.Id).Updates(map[string]interface{}{"status": 0, "department_id": 3}).RowsAffected
	return resInfo, int(count), nil
}

//批量添加数据库载体（导入）
func CreateStoreMore(url string) ([]model.HostLogs, error, int) {
	var ips = ""
	var names = ""
	var number = 0
	var str = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	var a []model.StoreAll
	var b []model.Res
	util.DbConn.Table("z_res left JOIN z_res_store ON z_res.id = z_res_store.res_id").Select("z_res.*,z_res_store.*").Find(&b).Order("Id desc").Where("z_res.status != 0")
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
	var store = make([]model.Store, 0)
	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
		return hostLogs, err, 0
	}
	var resSmall = new(model.Res)
	var StoreSmall = new(model.Store)
	rows, err := f.GetRows("存储设备模板")
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
			if rows[i][8] == "" {
				can = false
				logs.ErrorLog += "数据库名不能为空 "
			}
			if rows[i][9] == "" {
				can = false
				logs.ErrorLog += "物理主机不能为空 "
			}
			if rows[i][10] == "" {
				can = false
				logs.ErrorLog += "存储类型不能为空 "
			} else {
				if strings.Index(rows[i][11], "Oracle") > -1 {
					if rows[i][12] == "" {
						can = false
						logs.ErrorLog += "存储实例不能为空 "
					}
				}
			}
			if rows[i][11] == "" {
				can = false
				logs.ErrorLog += "存储版本不能为空 "
			}
			if rows[i][13] == "" {
				can = false
				logs.ErrorLog += "存储设备连接url不能为空 "
			}
			if rows[i][14] == "" {
				can = false
				logs.ErrorLog += "存储端口不能为空 "
			}
			if rows[i][15] == "" {
				can = false
				logs.ErrorLog += "存储设备驱动类名不能为空 "
			}
			if rows[i][16] == "" {
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
			StoreSmall.ResId = number
			StoreSmall.StoreClass = rows[i][15]
			StoreSmall.StoreName = rows[i][8]
			StoreSmall.StoreSid = rows[i][12]
			StoreSmall.StoreUrl = rows[i][13]
			StoreSmall.StorePort, _ = strconv.Atoi(rows[i][14])
			StoreSmall.StoreVersion, _ = strconv.Atoi(Substr(rows[i][11], strings.Index(rows[i][11], "_")))
			StoreSmall.StoreType, _ = strconv.Atoi(Substr(rows[i][10], strings.Index(rows[i][10], "_")))
			StoreSmall.ResHostId, _ = strconv.Atoi(Substr(rows[i][9], strings.Index(rows[i][9], "_")))
			resSmall.DicManufactorId, _ = strconv.Atoi(Substr(rows[i][7], strings.Index(rows[i][7], "_")))
			resSmall.DicCityId, _ = strconv.Atoi(Substr(rows[i][16], strings.Index(rows[i][16], "_")))
			resSmall.Id = number
			resSmall.DicTypeId = 1302
			store = append(store, *StoreSmall)
			res = append(res, *resSmall)
			fmt.Println("res2   ", res)
			logs.ErrorLog += "无异常 "
		}
		hostLogs = append(hostLogs, *logs)
	}
	countMain := util.DbConn.Table("z_res").Create(&res).RowsAffected
	count := util.DbConn.Table("z_res_store").Create(&store).RowsAffected
	fmt.Println(count)
	return hostLogs, err, int(countMain)
}
func Substr(str string, end int) string {
	rs := []rune(str)
	return string(rs[:end])
}

/*
导出存储设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportStore(resInfo model.StoreInfo) ([]model.ExportStore, error, int) {
	var count int64
	var exportList []model.ExportStore
	db := util.DbConn.Table("z_res res left join z_res_store sec on sec.res_id=res.id left join z_dictionary dicType on dicType.id=res.dic_type_id left join z_dictionary manufactor on manufactor.id=res.dic_manufactor_id  left join z_dictionary city on city.id=res.dic_city_id  left join z_department department on department.id=res.department_id left join z_res_account account on account.res_id=res.id  left join z_password_policy policy on policy.id=res.res_password_policy_id  left join z_res host on host.id=sec.res_host_id left join z_dictionary storeType on storeType.id=sec.store_type left join z_dictionary storeVersion on storeVersion.id=sec.store_version").Select("res.res_name,res.res_ip_v4,res.res_ip_v6,dicType.dic_name as res_type_name,department.dep_name as dep_name,account.res_account,account.res_account_name,policy.policy_name as res_password_policy,if(res.dic_connect_id=1,'是','否') as dic_connect_id,manufactor.dic_name as manufactor_name,sec.store_name,host.res_name as res_host_name,storeType.dic_name as store_type,storeVersion.dic_name as store_version,sec.store_sid,sec.store_url,sec.store_port,sec.store_class,city.dic_name as dic_city")
	db = db.Where(" res.dic_type_id = ?", 1302)
	db = db.Where(" res.status != ?", 0)
	if resInfo.Res.ResIpV4 != "" {
		db = db.Where("res.res_ip_v4 = ?", resInfo.Res.ResIpV4)
	}
	if resInfo.Res.ResName != "" {
		db = db.Where(" res.res_name like (?)", "%"+resInfo.Res.ResName+"%")
	}
	if resInfo.Store.StoreType != 0 {
		db = db.Where(" sec.store_type = ?", resInfo.Store.StoreType)
	}
	if resInfo.Res.DepartmentId != 0 {
		db = db.Where(" res.department_id = ?", resInfo.Res.DepartmentId)
	}
	db.Count(&count).Order("res.id desc").Find(&exportList)
	return exportList, nil, int(count)
}
