package service

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/util/cusfun"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:34
 */

/*
查询部门列表
*/
func DepartmentFindList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {

	res, count,err := dao.DepFindList(pojo)
	return res, err, count
}

/*
递归查询部门列表
 */
func DepTreeList(supId int) (rst interface{},err error,count int) {

	res,count,err := dao.GetTreeList(supId)
	return res,nil,count
}

/*
通过当前部门id查询下级部门
 */
func DepFindByIdList(id int) (rst interface{},err error,count int){

	res,count,err := dao.DepFindById(id)
	return res,nil,count

}

/*
部门新增
*/

func DepInsert(zDepartment model.Department) (department model.Department, err error, count int) {

	res,count,err := dao.DepInsert(zDepartment)
	return res, err, count

}

/*
部门更新
*/
func DepartmentUpdate(zDepartment model.Department) (department model.Department, err error, count int) {

	res, count, err := dao.DepUpdate(zDepartment)
	return res, err, count
}

/*
部门删除
*/
func DepartmentDel(zDepartment model.Department) (department model.Department, err error, count int) {

	res,count,err := dao.DepDel(zDepartment)
	return res, err, count
}

