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
 * @Date: 2021/4/3
 */
//分页查询列表
func ListHostProtocol(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {
	res, count, err := dao.ListHostProtocol(pojo)

	return res, nil, count
}

//添加协议管理
func CreateHostProtocol(resId int, pros []model.HostProtocol) (err error) {
	err = dao.CreateHostProtocol(resId, pros)
	return
}
