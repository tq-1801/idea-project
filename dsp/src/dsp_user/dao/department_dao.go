package dao

import (
	"dsp/src/dsp_user/model"
	"dsp/src/util"
	"dsp/src/util/cusfun"
	"errors"
	"strconv"
)

/**
 * @author  tianqiang
 * @date  2021/7/13 17:35
 */

/*
查询部门列表
*/
func DepFindList(pojo cusfun.ParamsPOJO) ( []model.Department,int, error) {
	var count int64
	var depFindList []model.Department
	db := util.DbConn.Model(&model.Department{})
	db = cusfun.GetSqlByParams(db, pojo, &count).Find(&depFindList)
	return depFindList, int(count), nil
}


type TreeList struct {
	Key              int         `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"key"`
	Title            string      `json:"title"`
	DepNumber        string      `json:"depNumber"`
	DicCityId        int         `json:"dicCityId"`
	ManagerUserId    string      `json:"managerUserId"`
	IsLeaf           int         `json:"isLeaf"`
	ZDesc            string      `json:"zDesc"`
	IsInlay          int         `json:"isInlay"`
	CreateDate       string      `json:"createDate"`
	CreateUserId     string      `json:"createUserId"`
	ModifyDate       string      `json:"modifyDate"`
	ModifyUserId     string      `json:"modifyUserId"`
	SupId            int         `json:"supId"`
	TreePath         string      `json:"treePath"`
	PasswordPolicyId int         `json:"passwordPolicyId"`
	TerminalPolicyId int         `json:"terminalPolicyId"`
	Children         []*TreeList `json:"children"`
}

/*
递归获取树形部门菜单
*/
func GetTreeList(supId int) ([]*TreeList, int, error) {
	var count int64
	var zDepartment []model.Department
	util.DbConn.Table("z_department").Where("sup_id = ?  ", supId).Find(&zDepartment).Count(&count) // 拿到所有父菜单(pid==0)
	treeList := []*TreeList{}
	for _, v := range zDepartment { // 循环所有父级部门
		child, _, _ := GetTreeList(v.Id) // 拿到每个父部门的子部门
		node := &TreeList{               // 拼装父级部门数据
			Key:              v.Id,
			Title:            v.DepName,
			DicCityId:        v.DicCityId,
			ManagerUserId:    v.ManagerUserId,
			IsLeaf:           v.IsLeaf,
			ZDesc:            v.ZDesc,
			CreateDate:       v.CreateDate,
			CreateUserId:     v.CreateUserId,
			ModifyDate:       v.ModifyDate,
			ModifyUserId:     v.ManagerUserId,
			SupId:            v.SupId,
			TreePath:         v.TreePath,
			PasswordPolicyId: v.PasswordPolicyId,
			TerminalPolicyId: v.TerminalPolicyId,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList, int(count), nil
}

//通过当前部门id查询下级部门
func DepFindById(id int) ([]model.Department,int,error)  {
	var count int64
	var DepFindById []model.Department

	util.DbConn.Table("z_department").Where("sup_id = ?",id).Order("id asc").Find(&DepFindById).Count(&count)

	return DepFindById,int(count),nil
}

/*
部门新增
*/
func DepInsert(zDepartment model.Department) (  model.Department, int, error) {

	var department model.Department
	if zDepartment.DepName == "" {
		return zDepartment,0,errors.New("部门名称不能为空")
	}
	//查询部门名称是否重复
	util.DbConn.Table("z_department").Where("dep_name = ? ", zDepartment.DepName).Take(&department)
	if department.DepName == zDepartment.DepName {
		return zDepartment,0,errors.New("部门名称已存在")
	}
	//查询部门编号是否重复
	util.DbConn.Table("z_department").Where("dep_number = ? ", zDepartment.DepNumber).Take(&department)
	if zDepartment.DepNumber != "" && department.DepNumber == zDepartment.DepNumber {
		return zDepartment,0,errors.New("部门编号已存在")
	}
	zDepartment.IsLeaf = 1
	zDepartment.IsInlay = 0
	//定义父节点的路径
	if zDepartment.TreePath == "" {
		//找出上级的父节点路径
		util.DbConn.Table("z_department").Where("id = ?", zDepartment.SupId).Take(&department)

		//把父节点int类型转换为string 类型在加上上级部门的父节点路径
		zDepartment.TreePath = strconv.Itoa(zDepartment.SupId) + "," + department.TreePath
	}

	count := util.DbConn.Create(&zDepartment).RowsAffected
	//修改已增加的部门的上级部门的子节点为0
	util.DbConn.Table("z_department").Where("id = ?", zDepartment.SupId).Update("is_leaf", 0)

	return zDepartment, int(count), nil

}

/*
部门更新
*/
func DepUpdate(zDepartment model.Department) (model.Department, int, error) {

	var department model.Department
	var all model.Department

	//通过部门id查询修改前的部门信息
	util.DbConn.Table("z_department").Where("id = ? ", zDepartment.Id).First(&all)
	if zDepartment.DepName == "" {
		return zDepartment,1,errors.New("部门名称不能为空")
	}
	//判断部门名称是否重复
	util.DbConn.Table("z_department").Where("dep_name = ? ", zDepartment.DepName).Take(&department)
	if all.DepName != zDepartment.DepName && department.DepName == zDepartment.DepName {
		return zDepartment,1,errors.New("部门名称已存在")
	}
	//判断部门编号是否重复
	util.DbConn.Table("z_department").Where("dep_number = ? ", zDepartment.DepNumber).Take(&department)
	if all.DepNumber != zDepartment.DepNumber && department.DepNumber == zDepartment.DepNumber {
		return zDepartment,1,errors.New("部门编号已存在")
	}
	//定义父节点的路径
	if zDepartment.TreePath == "" {
		//查询上级部门的父节点路径
		util.DbConn.Table("z_department").Where("id = ?", zDepartment.SupId).Take(&department)
		//把父节点int类型转换为string 类型在加上上级部门的父节点路径
		zDepartment.TreePath = strconv.Itoa(zDepartment.SupId) + "," + department.TreePath
	}

	count := util.DbConn.Save(&zDepartment).RowsAffected
	return zDepartment, int(count), nil
}

/*
部门删除
*/
func DepDel(zDepartment model.Department) (model.Department,int, error) {

	var all model.Department
	var department model.Department
	var user model.User
	var res model.Res

	//通过部门id查询删除前的部门信息
	util.DbConn.Table("z_department").Where("id = ? ", zDepartment.Id).First(&all)
	//初始化数据不允许删除
	if all.IsInlay == 1 {
		return all,1, errors.New("初始化数据不可被删除")
	}
	//通过部门id查询部门下的子部门
	util.DbConn.Table("z_department").Where("sup_id = ? ", zDepartment.Id).Take(&department)
	//部门下有子部门不能被删除
	if zDepartment.Id == department.Id {
		return all,1, errors.New("当前部门下有子部门，需删除子部门后才能删除当前部门")
	}
	//查询部门下是否有用户
	util.DbConn.Table("z_user").Where("department_id = ? and status != 0 ", zDepartment.Id).Take(&user)
	//部门下有用户不能被删除
	if zDepartment.Id == user.DepartmentId {
		return all,1, errors.New("当前部门下有用户，不可被删除")
	}
	//查询部门下是否有资源
	util.DbConn.Table("z_res").Where("department_id = ? and status != 0", zDepartment.Id).Take(&res)
	//部门下有资源不能被删除
	if zDepartment.Id == res.DepartmentId {
		return all,1, errors.New("当前部门下有资源，不可被删除")
	}
	if DepartmentSupIdList(all.SupId) == 1 {
		//没有同级部门时，修改上级部门子节点为1之后在删除
		util.DbConn.Table("z_department").Where("id = ?", all.SupId).Update("is_leaf", 1)
		//删除
		count := util.DbConn.Delete(&zDepartment).RowsAffected
		return zDepartment,int(count),nil
	}
	//当有同级部门时，不做任何修改,直接删除
	if DepartmentSupIdList(all.SupId) > 1 {
		count := util.DbConn.Delete(&zDepartment).RowsAffected
		return zDepartment,int(count),nil
	}
	return zDepartment, 1, nil
}


//当执行删除操作时，查询同级别是否还有部门
func DepartmentSupIdList(supId int) (count int64) {
	util.DbConn.Table("z_department").Where("sup_id = ?", supId).Count(&count)
	return count
}


