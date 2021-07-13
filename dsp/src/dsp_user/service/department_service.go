package service

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/util/cusfun"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:34
 */
/*
查询部门
*/
func DepartmentFindList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {

	res, err, count := dao.DepartmentFindList(pojo)
	return res, err, count
}