package service

import (
	"dsp/src/dsp_resource/model"
	"dsp/src/util/cusfun"
)

/**
 * @Author: zhaojia
 * @Description:业务系统信息管理
 * @Version: 1.0.0
 * @Date: 2021/3/25
 */
/*
查询业务系统信息管理
*/
func ResBusinessList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, err, count := dao.ResBusinessList(pojo)

	return res, err, count
}

//添加业务系统信息管理
func CreateResBusiness(resInfo model.ResBusinessInfo) (model.ResBusinessInfo, int, error) {
	res, count, err := dao.CreateResBusiness(resInfo)
	return res, count, err
}

//修改业务系统信息管理
func UpdateResBusiness(resInfo model.ResBusinessInfo) (model.ResBusinessInfo, int, error) {
	res, count, err := dao.UpdateResBusiness(resInfo)
	return res, count, err
}

//删除业务系统信息管理
func DeleteResBusiness(resInfo model.ResBusinessInfo) (model.ResBusinessInfo, int, error) {
	res, count, err := dao.DeleteResBusiness(resInfo)
	return res, count, err
}

/*
导出业务系统列表
*/
func ExportResBusiness(resInfo model.ResBusinessInfo) ([]model.ExportBusiness, error, int) {
	res, err, count := dao.ExportResBusiness(resInfo)
	return res, err, count
}

//导入业务系统
func CreateBusinessMore(filePath string) ([]model.ImportErrorLogs, error, int) {
	res, err, count := dao.CreateBusinessMore(filePath)
	return res, err, count
}
