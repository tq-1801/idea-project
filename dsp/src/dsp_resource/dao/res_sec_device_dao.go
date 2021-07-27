package dao

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"fmt"
	"regexp"

	"strconv"
	"strings"
)

/**
 * @Author: zhaojia
 * @Description:安全设备管理
 * @Version: 1.0.0
 * @Date: 2021/3/22
 */
/*
查询安全设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ResSecDeviceList(pojo cusfun.ParamsPOJO) (res []model.ResSecDevicePage, err error, total int) {
	var count int64
	db := util.DbConn.Table("(select res.id,res.res_name,res.res_ip_v4,res.res_ip_v6,res.dic_connect_id,res.dic_type_id,res.dic_city_id,res.dic_manufactor_id,res.department_id,res.res_manager_user_id,res.`status`,res.z_desc,date_format(res.create_date, '%Y-%m-%d %H:%i:%s') as create_date,res.create_user_id,date_format(res.modify_date, '%Y-%m-%d %H:%i:%s') as modify_date,res.modify_user_id,res.res_password_policy_id,res.secure_port,res.somf_ip,sec.res_id,sec.sec_device_type,sec.admin_account,sec.admin_password,sec.root_password,sec.admin_con_sym,sec.root_con_sym,sec.res_skip_id,sec.web_url,sec.protocol_type,sec.protocol_port,sec.dic_version_id,sec.dic_model_id,subtype.dic_name as sec_device_type_name,department.dep_name dep_name,policy.policy_name policy_name,skip.res_name skip_name,USER.uname from z_res res left join z_res_sec_device sec on sec.res_id=res.id left join z_dictionary subtype on subtype.id=sec.sec_device_type LEFT JOIN z_department department ON department.id = res.department_id LEFT JOIN z_password_policy policy ON policy.id = res.res_password_policy_id LEFT JOIN z_res skip ON skip.id = sec.res_skip_id LEFT JOIN z_user USER ON USER.uid = res.res_manager_user_id) t").Select("id,res_name,res_ip_v4,res_ip_v6,dic_connect_id,dic_type_id,dic_city_id,dic_manufactor_id,department_id,res_manager_user_id,`status`,z_desc,create_date,create_user_id,modify_date,modify_user_id,res_password_policy_id,secure_port,somf_ip,res_id,sec_device_type,admin_account,admin_password,root_password,admin_con_sym,root_con_sym,res_skip_id,web_url,protocol_type,protocol_port,dic_version_id,dic_model_id,sec_device_type_name,dep_name,policy_name,skip_name,uname ")
	db = db.Where("dic_type_id = ?", 1304)
	db = db.Where("status != 0")
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}

//添加安全设备
func CreateResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	var list = []string{"SSH", "TELNET"}
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	resInfo.Res.DicTypeId = 1304
	resInfo.Res.Status = 1
	errMain := tx.Create(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.ResSecDevice.ResId = int64(resInfo.Res.Id)
	//保存子表信息
	err := tx.Create(&resInfo.ResSecDevice).Error

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
	if errMain == nil && err == nil && errpro == nil {
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

//修改安全设备
func UpdateResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	var count int64
	//开始事务
	tx := util.DbConn.Begin()
	//保存主表信息
	errMain := tx.Save(&resInfo.Res).Error
	//将添加的主键赋值给子表
	resInfo.ResSecDevice.ResId = int64(resInfo.Res.Id)
	//保存子表信息
	err := tx.Save(&resInfo.ResSecDevice).Error

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

//删除安全设备，子表联代删除
func DeleteResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	count := util.DbConn.Table("z_res").Where("id=?", resInfo.Res.Id).Updates(map[string]interface{}{"status": 0, "department_id": 3}).RowsAffected
	return resInfo, int(count), nil
}

/*
导出安全设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportResSecDevice(resInfo model.ResSecDeviceInfo) ([]model.ExportSecDevice, error, int) {
	var count int64
	var exportList []model.ExportSecDevice
	db := util.DbConn.Table("z_res res left join z_res_sec_device sec on sec.res_id=res.id left join z_dictionary dicType on dicType.id=res.dic_type_id left join z_dictionary manufactor on manufactor.id=res.dic_manufactor_id left join z_dictionary city on city.id=res.dic_city_id left join z_department department on department.id=res.department_id left join z_res_account account on account.res_id=res.id left join z_password_policy policy on policy.id=res.res_password_policy_id left join z_dictionary subtype on subtype.id=sec.sec_device_type left join z_dictionary protocol on protocol.id=sec.protocol_type left join z_dictionary version on version.id=sec.dic_version_id left join z_dictionary model on model.id=sec.dic_model_id").Select(" res.res_name,res.res_ip_v4,res.res_ip_v6,dicType.dic_name as res_type_name,department.dep_name as dep_name,account.res_account,account.res_account_name,policy.policy_name as res_password_policy,if(res.dic_connect_id=1,'是','否') as dic_connect_id,manufactor.dic_name as manufactor_name,subtype.dic_name as sec_device_type,'' as res_skip,protocol.dic_name as protocol_type,sec.protocol_port,version.dic_name as dic_version,model.dic_name as dic_model,city.dic_name as dic_city")
	db = db.Where(" res.dic_type_id = ?", 1304)
	db = db.Where(" res.status != ?", 0)
	if resInfo.Res.ResIpV4 != "" {
		db = db.Where("res.res_ip_v4 = ?", resInfo.Res.ResIpV4)
	}
	if resInfo.Res.ResName != "" {
		db = db.Where(" res.res_name like (?)", "%"+resInfo.Res.ResName+"%")
	}
	if resInfo.ResSecDevice.SecDeviceType != 0 {
		db = db.Where(" sec.sec_device_type = ?", resInfo.ResSecDevice.SecDeviceType)
	}
	if resInfo.Res.DepartmentId != 0 {
		db = db.Where(" res.department_id = ?", resInfo.Res.DepartmentId)
	}
	db.Count(&count).Order("res.id desc").Find(&exportList)
	return exportList, nil, int(count)
}

//导入安全设备
func CreateDeviceMore(url string) ([]model.ImportErrorLogs, error, int) {
	var deviceLogs = make([]model.ImportErrorLogs, 0)
	number := 0
	//取出主表最大值，赋值给number，用于生成资产的主键
	util.DbConn.Table("z_res").Select("max(id)").Find(&number)
	//取出所有的设备名称和IP
	var ips = ""
	var names = ""
	var allDevice []model.Res
	util.DbConn.Table("z_res").Select("z_res.*").Find(&allDevice).Order("id desc")
	if len(allDevice) > 0 {
		for i := 0; i < len(allDevice); i++ {
			ips = ips + allDevice[i].ResIpV4 + "@"
			names = names + allDevice[i].ResName + "@"
		}
	}

	f, err := excelize.OpenFile(url)
	if err != nil {
		fmt.Println(err)
		return deviceLogs, err, 0
	}

	var resOne = new(model.Res)
	var deviceOne = new(model.ResSecDevice)

	var res = make([]model.Res, 0)
	var device = make([]model.ResSecDevice, 0)

	rows, err := f.GetRows("安全设备模板")
	//从1开始  第一行不要,for结束以后rows就是校验完的
	for i := 1; i < len(rows); i++ {
		can := true
		var logs = new(model.ImportErrorLogs)
		logs.Name = rows[i][0]
		logs.SecondMessage = rows[i][1]
		//必有字段有无校验
		if len(rows[i]) == 0 {
			can = false
			logs.ErrorLog += "空行 "
		} else {
			if rows[i][0] == "" {
				can = false
				logs.ErrorLog += "资产名称不能为空 "
			} else {
				if strings.Index(names, rows[i][0]) > -1 {
					can = false
					logs.ErrorLog += "资产名称重复 "
				} else {
					names = names + rows[i][0] + "@"
				}
			}
			if rows[i][1] == "" { //校验IP是否合法
				can = false
				logs.ErrorLog += "资产IP不能为空 "
			} else {
				regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
				match, _ := regexp.MatchString(regStr, rows[i][1])
				if !match {
					can = false
					logs.ErrorLog += "资产IP格式错误 "
				} else {
					if strings.Index(ips, rows[i][1]) > -1 {
						can = false
						logs.ErrorLog += "IP重复 "
					} else {
						ips = ips + rows[i][1] + "@"
					}
				}
			}
			if rows[i][3] == "" {
				can = false
				logs.ErrorLog += "部门名称不能为空 "
			}
			if rows[i][6] == "" {
				can = false
				logs.ErrorLog += "是否可连接不能为空 "
			}
			if rows[i][8] == "" {
				can = false
				logs.ErrorLog += "安全设备类型不能为空 "
			}
			if rows[i][9] == "" {
				can = false
				logs.ErrorLog += "协议类型不能为空 "
			}
			if rows[i][10] == "" {
				can = false
				logs.ErrorLog += "协议端口不能为空 "
			}
		}
		if can {
			number++
			resOne.Id = number
			resOne.DepartmentId, _ = strconv.Atoi(Substr(rows[i][3], strings.Index(rows[i][3], "_")))
			resOne.ResPasswordPolicyId, _ = strconv.Atoi(Substr(rows[i][5], strings.Index(rows[i][5], "_")))
			resOne.ResName = rows[i][0]
			resOne.ResIpV4 = rows[i][1]
			resOne.ResIpV6 = rows[i][2]
			resOne.ResManagerUserId = rows[i][4]
			resOne.DicTypeId = 1304
			resOne.Status = 1
			if rows[i][6] == "是" {
				resOne.DicConnectId = 1
			} else {
				resOne.DicConnectId = 0
			}
			resOne.DicManufactorId, _ = strconv.Atoi(Substr(rows[i][7], strings.Index(rows[i][7], "_")))
			if len(rows[i]) == 12 {
				resOne.DicCityId, _ = strconv.Atoi(Substr(rows[i][11], strings.Index(rows[i][11], "_")))
			}
			deviceOne.ResId = int64(number)
			deviceOne.SecDeviceType, _ = strconv.ParseInt(Substr(rows[i][8], strings.Index(rows[i][8], "_")), 10, 64)
			deviceOne.ProtocolType, _ = strconv.ParseInt(Substr(rows[i][9], strings.Index(rows[i][9], "_")), 10, 64)
			deviceOne.ProtocolPort, _ = strconv.ParseInt(rows[i][10], 10, 64)
			device = append(device, *deviceOne)
			res = append(res, *resOne)
			logs.ErrorLog += "无异常 "
		}
		if logs.ErrorLog != "" {
			deviceLogs = append(deviceLogs, *logs)
		}
	}
	//保存主表信息
	countMain := util.DbConn.Table("z_res").Create(&res).RowsAffected
	//保存子表信息
	count := util.DbConn.Table("z_res_sec_device").Create(&device).RowsAffected
	fmt.Println(count)
	return deviceLogs, err, int(countMain)
}
