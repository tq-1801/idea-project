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
/*
列表查询
*/
func ResSecDeviceList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, err, count := dao.ResSecDeviceList(pojo)

	return res, err, count
}

//添加安全设备
func CreateResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	res, count, err := dao.CreateResSecDevice(resInfo)
	return res, count, err
}

//修改安全设备
func UpdateResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	res, count, err := dao.UpdateResSecDevice(resInfo)
	return res, count, err
}

//删除安全设备
func DeleteResSecDevice(resInfo model.ResSecDeviceInfo) (model.ResSecDeviceInfo, int, error) {
	res, count, err := dao.DeleteResSecDevice(resInfo)
	return res, count, err
}

/*
导出安全设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportResSecDevice(resInfo model.ResSecDeviceInfo) ([]model.ExportSecDevice, error, int) {
	res, err, count := dao.ExportResSecDevice(resInfo)
	return res, err, count
}

//导入安全设备
func CreateDeviceMore(filePath string) ([]model.ImportErrorLogs, error, int) {
	res, err, count := dao.CreateDeviceMore(filePath)
	return res, err, count
}
