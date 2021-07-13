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
	Id               int    `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"id"`
	DepName          string `json:"depName"`
	DepNumber        string `json:"depNumber"`
	DicCityId        int    `json:"dicCityId"`
	ManagerUserId    string `json:"managerUserId"`
	IsLeaf           int    `json:"isLeaf"`
	ZDesc            string `json:"zDesc"`
	IsInlay          int    `json:"isInlay"`
	CreateDate       string `json:"createDate"`
	CreateUserId     string `json:"createUserId"`
	ModifyDate       string `json:"modifyDate"`
	ModifyUserId     string `json:"modifyUserId"`
	SupId            int    `json:"supId"`
	TreePath         string `json:"treePath"`
	PasswordPolicyId int    `json:"passwordPolicyId"`
	TerminalPolicyId int    `json:"terminalPolicyId"`
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
