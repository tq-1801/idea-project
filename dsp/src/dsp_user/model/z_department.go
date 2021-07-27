package model

import (
	"gorm.io/gorm"
	"time"
)

/**

* @author tianqiang
* @date 2021/7/13.

 */

type Department struct {
	Id               int    `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id" mapstructure:"id" validate:"" trans:"部门id"`
	DepName          string `json:"depName" mapstructure:"depName" validate:"require" trans:"部门"`
	DepNumber        string `json:"depNumber" mapstructure:"depNumber" validate:"" trans:"部门编号"`
	DicCityId        int    `json:"dicCityId" mapstructure:"dicCityId" validate:"" trans:"所属地市"`
	ManagerUserId    string `json:"managerUserId" mapstructure:"managerUserId" validate:"" trans:"部门负责人账号"`
	IsLeaf           int    `json:"isLeaf" mapstructure:"isLeaf" validate:"" trans:"子节点"`
	ZDesc            string `json:"zDesc" mapstructure:"zDesc" validate:"" trans:"描述"`
	IsInlay          int    `json:"isInlay" mapstructure:"isInlay" validate:"" trans:"初始化数据"`
	CreateDate       string `json:"createDate" mapstructure:"createDate" validate:"" trans:"数据创建时间"`
	CreateUserId     string `json:"createUserId" mapstructure:"createUserId" validate:"" trans:"数据修改用户账号"`
	ModifyDate       string `json:"modifyDate" mapstructure:"modifyDate" validate:"" trans:"数据修改时间"`
	ModifyUserId     string `json:"modifyUserId" mapstructure:"modifyUserId" validate:"" trans:"数据修改账号"`
	SupId            int    `json:"supId" mapstructure:"supId" validate:"" trans:"父节点id"`
	TreePath         string `json:"treePath" mapstructure:"treePath" validate:"" trans:"父节点的路径"`
	PasswordPolicyId int    `json:"passwordPolicyId" mapstructure:"passwordPolicyId" validate:"" trans:"密码策略id"`
	TerminalPolicyId int    `json:"terminalPolicyId" mapstructure:"terminalPolicyId" validate:"" trans:"终端安全策略id"`
}
type TreeList struct {
	Id               int        `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id"`
	DepName          string     `json:"depName"`
	DepNumber        string     `json:"depNumber"`
	DicCityId        int        `json:"dicCityId"`
	ManagerUserId    string     `json:"managerUserId"`
	IsLeaf           int        `json:"isLeaf"`
	ZDesc            string     `json:"zDesc"`
	IsInlay          int        `json:"isInlay"`
	CreateDate       string     `json:"createDate"`
	CreateUserId     string     `json:"createUserId"`
	ModifyDate       string     `json:"modifyDate"`
	ModifyUserId     string     `json:"modifyUserId"`
	SupId            int        `json:"supId"`
	TreePath         string     `json:"treePath"`
	PasswordPolicyId int        `json:"passwordPolicyId"`
	TerminalPolicyId int        `json:"terminalPolicyId"`
	Children         []TreeList `json:"children"`
}
type DataRes struct {
	Data []*TreeList `json:"data"`
}

func (Department) TableName() string {
	return "z_department"
}
func (zDepartment *Department) BeforeCreate(tx *gorm.DB) (err error) {
	zDepartment.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	zDepartment.ModifyDate = time.Now().Format("0000-00-00 00:00:00")
	return
}

func (zDepartment *Department) BeforeSave(tx *gorm.DB) (err error) {
	zDepartment.ModifyDate = time.Now().Format("2006-01-02 15:04:05")
	//zDepartment.CreateDate = time.Now().Format("0000-00-00 00:00:00")
	return
}
