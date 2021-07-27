package service

import (
	"dsp/src/dsp_resource/dao"
	"dsp/src/dsp_resource/model"
	"dsp/src/util/cusfun"
)

/**
 * @Author: wangyue
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/3/23
 */

func HostList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, err, count := dao.HostList(pojo)
	return res, err, count
}

func CreateHost(client model.HostInfo) (rst interface{}, err error, count int) {
	res, err, count := dao.CreateHost(client)
	return res, err, count
}

func UpdateHost(client model.HostInfo) (rst interface{}, count int, err error) {
	res, count, err := dao.UpdateHost(client)
	return res, count, err
}

func DeleteHost(client model.HostInfo) (rst interface{}, count int, err error) {
	res, count, err := dao.DeleteHost(client)
	return res, count, err
}

func CreateHostMore(url string) (rst interface{}, err error, count int) {
	res, err, count := dao.CreateHostMore(url)
	return res, err, count
}

/*
导出主机设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportHost(resInfo model.HostInfo) ([]model.ExportHost, error, int) {
	res, err, count := dao.ExportHost(resInfo)
	return res, err, count
}

//测试用户名和密码是否可以登录
func LoginHost(params map[string]interface{}) (err error) {
	err = dao.LoginHost(params)
	return
}
