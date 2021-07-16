/**
 * @Author: yanghang
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/3/31
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id           int    `gorm:"primary_key" json:"id"`
	Rolename     string `json:"rolename"`
	ZDesc        string `json:"zDesc"`
	Isinlay      string `json:"isinlay"`
	Createdate   string `json:"createdate"`
	Createuserid string `json:"createuserid"`
	Modifydate   string `json:"modifydate"`
	Modifyuserid string `json:"modifyuserid"`
}

func (role *Role) BeforeCreate(tx *gorm.DB) (err error) {
	role.Createdate = time.Now().Format("2006-01-02 15:04:05")
	role.Modifydate = time.Now().Format("0000-00-00 00:00:00")
	return
}
func (role *Role) BeforeSave(tx *gorm.DB) (err error) {
	role.Modifydate = time.Now().Format("2006-01-02 15:04:05")
	return
}

func (Role) TableName() string {
	return "z_role"
}
