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

//网络设备列表查询

func ResNetworkDeviceList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, err, count := dao.ResNetworkDeviceList(pojo)

	return res, err, count
}

//添加网络设备
func CreateResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	res, count, err := dao.CreateResNetworkDevice(resInfo)
	return res, count, err
}

//修改网络设备
func UpdateResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	res, count, err := dao.UpdateResNetworkDevice(resInfo)
	return res, count, err
}

//删除网络设备
func DeleteResNetworkDevice(resInfo model.ResNetworkDeviceInfo) (model.ResNetworkDeviceInfo, int, error) {
	res, count, err := dao.DeleteResNetworkDevice(resInfo)
	return res, count, err
}
func CreateNetworkMore(url string) (rst interface{}, err error, count int) {
	res, err, count := dao.CreateNetworkMore(url)
	return res, err, count
}

/*
导出网络设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportDevice(resInfo model.ResNetworkDeviceInfo) ([]model.ExportNetwork, error, int) {
	res, err, count := dao.ExportDevice(resInfo)
	return res, err, count
}
