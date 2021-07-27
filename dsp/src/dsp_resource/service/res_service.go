package service

import (
	"dsp/src/dsp_resource/dao"
	"dsp/src/util/cusfun"
)

func TelnetRes(params map[string]interface{}) bool {
	bool := dao.TelnetRes(params)
	return bool
}

//通过部门id获取部门名称
func GetDeptName(deptId int) string {
	deptName := dao.GetDeptName(deptId)
	return deptName
}

//判断资产的IP或者名称是否存在
func ResIsExists(pojo cusfun.ParamsPOJO) (err error, total int) {
	err, total = dao.ResIsExists(pojo)
	return
}

func ResStoreIsExists(ip, name, storetype string) (err error, total int) {
	err, total = dao.ResStoreIsExists(ip, name, storetype)
	return
}

func ResAccountIsExists(resId, resAccount, accountId string) (err error, total int) {
	err, total = dao.ResAccountIsExists(resId, resAccount, accountId)
	return
}
