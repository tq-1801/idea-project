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

func StoreList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, err, count := dao.StoreList(pojo)
	return res, err, count
}

func CreateStore(client model.StoreInfo) (rst interface{}, err error, count int) {
	res, err, count := dao.CreateStore(client)
	return res, err, count
}

func UpdateStore(client model.StoreInfo) (rst interface{}, count int, err error) {
	res, count, err := dao.UpdateStore(client)
	return res, count, err
}

func DeleteStore(client model.StoreInfo) (rst interface{}, count int, err error) {
	res, count, err := dao.DeleteStore(client)
	return res, count, err
}
func CreateStoreMore(url string) (rst interface{}, err error, count int) {
	res, err, count := dao.CreateStoreMore(url)
	return res, err, count
}

/*
导出存储设备列表
可以根据资产名称、资产IP，部门，类型进行查询
*/
func ExportStore(resInfo model.StoreInfo) ([]model.ExportStore, error, int) {
	res, err, count := dao.ExportStore(resInfo)
	return res, err, count
}
