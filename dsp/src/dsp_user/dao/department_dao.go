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
查询部门列表
*/
func DepFindList(pojo cusfun.ParamsPOJO) (res []model.Department, err error, total int) {
	var count int64
	db := util.DbConn.Model(&model.Department{})
	err = cusfun.GetSqlByParams(db, pojo, &count).Find(&res).Error
	total = int(count)
	return
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
func DepInsert(zDepartment model.Department) (model.Department,error) {

	if zDepartment.ZDesc == "" {
		zDepartment.ZDesc = "无描述"
	}
	zDepartment.IsLeaf = 1
	zDepartment.IsInlay = 0
	err := util.DbConn.Create(&zDepartment).Error
	return zDepartment,err

}

//通过部门名称查询
func DepFindByName(name string) (count int64) {

	util.DbConn.Table("z_department").Where("dep_name = ? ", name).Count(&count)
	return count

}

/*
根据部门编号查询
*/
func DepFindByNumber(depNumber string) (count int64) {
	util.DbConn.Table("z_department").Where("dep_number = ? ", depNumber).Count(&count)
	return count
}

/*
通过父节点id，查询上级部门
*/
func DepFindBySupId(supId int) model.Department {

	var DepFindBySupId model.Department
	util.DbConn.Table("z_department").Where("id = ?", supId).Take(&DepFindBySupId)
	return DepFindBySupId
}

/*
修改已增加的部门的上级部门的子节点为0
*/
func DepIsLeafUpdate(supId int) ([]model.Department, int, error) {
	var count int
	var DepartmentIsLeafUpdate []model.Department

	util.DbConn.Table("z_department").Where("id = ?", supId).Update("is_leaf", 0)
	return DepartmentIsLeafUpdate, count, nil
}

/*
部门更新
*/
func DepartmentUpdate(zDepartment model.Department) (model.Department, int, error) {
	count := util.DbConn.Save(&zDepartment).RowsAffected
	return zDepartment, int(count), nil
}

/*
部门删除
*/
func DepDel(zDepartment model.Department) (model.Department, error) {
	err := util.DbConn.Delete(&zDepartment).Error
	return zDepartment, err
}

/*
查询部门下是否是初始化数据
*/
func DepartmentById(id int) (res model.Department) {
	util.DbConn.Table("z_department").Where("id = ? ", id).First(&res)
	return res
}

/*
查询部门下是否有子部门
*/
func DepartmentFindBySupId(id int) (count int64) {
	util.DbConn.Table("z_department").Where("sup_id = ? ", id).Count(&count)
	return count
}

/*
查询部门下是否有用户
*/
func DepartmentFindById(id int) (count int64) {
	util.DbConn.Table("z_user").Where("department_id = ? ", id).Count(&count)
	return count
}

/*
查询部门下是否有资产
*/
func DepartmentFindId(id int) (count int64) {
	util.DbConn.Table("z_res").Where("department_id = ? ", id).Count(&count)
	return count
}

/*
当执行删除操作时，查询同级别是否还有部门
*/
func DepartmentSupIdList(supId int) (count int64) {
	util.DbConn.Table("z_department").Where("sup_id = ?", supId).Count(&count)
	return count
}

/*
在删除后没有同级部门则修改已删除的部门的上级部门的子节点为1
*/
func DepartmentIsLeafSupIdUpdate(supId int) ([]model.Department, int, error) {
	var count int
	var DepartmentIsLeafSupIdUpdate []model.Department
	util.DbConn.Table("z_department").Where("id = ?", supId).Update("is_leaf", 1)
	return DepartmentIsLeafSupIdUpdate, count, nil
}
