package service

import (
	"dsp/src/dsp_resource/dao"
	"dsp/src/dsp_resource/model"
	"dsp/src/util/cusfun"
)

/**
 * @Author: zhaojia
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/3/22
 */
//分页查询列表
func ListAccount(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, count, err := dao.ListAccount(pojo)

	return res, nil, count
}

//添加账号
func CreateAccount(account model.Account) (rst interface{}, err error, count int) {
	res, count, err := dao.CreateAccount(account)

	return res, err, count
}

//修改账号
func UpdateAccount(params map[string]interface{}, account model.Account) (rst interface{}, err error, count int) {
	res, count, err := dao.UpdateAccount(params, account)

	return res, err, count
}

//删除账号
func DeleteAccount(account model.Account) (rst interface{}, err error, count int) {
	res, count, err := dao.DeleteAccount(account)

	return res, nil, count
}

/*
导出账号信息
*/
func ExportAccount(account model.Account) ([]model.ExportAccount, error, int) {
	res, err, count := dao.ExportAccount(account)
	return res, err, count
}

//导入安全设备
func CreateAccountMore(filePath string, resId int) ([]model.ImportErrorLogs, error, int) {
	res, err, count := dao.CreateAccountMore(filePath, resId)
	return res, err, count
}

func ListAccountByResId(resId string) (res []model.AccountPage, total int, err error) {
	res, count, err := dao.ListAccountByResId(resId)
	return res, count, nil
}
