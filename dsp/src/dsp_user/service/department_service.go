package service

import (
	"dsp/src/dsp_user/dao"
	"dsp/src/dsp_user/model"
	"dsp/src/util/cusfun"
	"errors"
	"strconv"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:34
 */

/*
查询部门列表
*/
func DepartmentFindList(pojo cusfun.ParamsPOJO) (rst interface{}, err error, count int) {

	res, err, count := dao.DepFindList(pojo)
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

	//查询部门名称是否重复
	if dao.DepFindByName(zDepartment.DepName) > 0 {
		err := errors.New("部门名称已存在")
		return zDepartment, err, 0
	}
	if zDepartment.DepNumber != "" {
		// 查询部门编号是否重复
		if dao.DepFindByNumber(zDepartment.DepNumber) > 0 {
			err := errors.New("部门编号已存在")
			return zDepartment, err, 0
		}
	}

	// 定义父节点的路径
	if zDepartment.TreePath == "" {
		//调用方法找出上级的父节点路径
		err := dao.DepFindBySupId(zDepartment.SupId).TreePath
		//把int类型转换为string类型
		a := strconv.Itoa(zDepartment.SupId)
		zDepartment.TreePath = a + "," + err
	}
	res,count,err := dao.DepInsert(zDepartment)
	dao.DepIsLeafUpdate(zDepartment.SupId)

	return res, err, count

}

/*
部门更新
*/
func DepartmentUpdate(zDepartment model.Department) (department model.Department, err error, count int) {

	res := dao.DepartmentById(zDepartment.Id)
	//查询部门名称是否重复
	if res.DepName != zDepartment.DepName {
		if dao.DepFindByName(zDepartment.DepName) > 0 {
			err := errors.New("部门名称已存在")
			return zDepartment, err, 0
		}
	}
	// 查询部门编号是否重复
	if res.DepNumber != zDepartment.DepNumber {
		if zDepartment.DepNumber != "" && dao.DepFindByNumber(zDepartment.DepNumber) > 0 {

				err := errors.New("部门编号已存在")
				return zDepartment, err, 0
			}
		}
	// 定义父节点的路径
	if zDepartment.TreePath == "" {
		//调用方法找出上级的父节点路径
		err := dao.DepFindBySupId(zDepartment.SupId).TreePath
		//把int类型转换为string类型
		a := strconv.Itoa(zDepartment.SupId)
		zDepartment.TreePath = a + "," + err
	}
	res, count, err = dao.DepUpdate(zDepartment)
	return res, nil, count
}

/*
部门删除
*/
func DepartmentDel(zDepartment model.Department) (department model.Department, err error, count int) {
	res := dao.DepartmentById(zDepartment.Id)
	//初始化数据不允许删除
	if res.IsInlay == 1 {
		err = errors.New("初始化数据不可被删除")
		return res, err, count
	}
	//部门下有子部门不能被删除
	if dao.DepartmentFindBySupId(zDepartment.Id) > 0 {
		err = errors.New("当前部门下有子部门，需删除子部门后才能删除当前部门")
		return res, err, count
	}
	//部门下有用户不能被删除
	if dao.DepartmentFindById(zDepartment.Id) > 0 {
		err = errors.New("当前部门下有用户，不可被删除")
		return res, err, count

	}
	//部门下有资源不能被删除
	if dao.DepartmentFindId(zDepartment.Id) > 0 {
		err = errors.New("当前部门下有资源，不可被删除")
		return res, err, count

	}
	//没有同级部门时，修改上级部门子节点为1之后在删除
	if dao.DepartmentSupIdList(res.SupId) == 1 {
		dao.DepartmentIsLeafSupIdUpdate(res.SupId)
		//删除
		res, count,err:= dao.DepDel(zDepartment)
		return res, err, count
	}
	//当有同级部门时，不做任何修改,直接删除
	if dao.DepartmentSupIdList(res.SupId) > 1 {
		//删除
		res,count,err := dao.DepDel(zDepartment)
		return res, err, count
	}
	return res, err, count
}

