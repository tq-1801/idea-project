package dao

import (
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:35
 */
/*
查询部门
*/
func DepartmentFindList(pojo cusfun.ParamsPOJO) (res []model.Department, err error, total int) {
	var count int64
	db := util.DbConn.Model(&model.Department{})
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
}