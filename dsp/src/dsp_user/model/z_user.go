/**
 * @Author: tianqiang
 * @Date: 2021/7/14
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uid                  string `gorm:"primary_key" json:"uid"`
	Uname                string `json:"uname"`
	Sex                  string `json:"sex"`
	Cardid               string `json:"cardid"`
	Birthdate            string `json:"birthdate"`
	Mail                 string `json:"mail"`
	Mobile               string `json:"mobile"`
	Password             string `json:"password"`
	Status               string `json:"status"`
	Lastlogin            string `json:"lastlogin"`
	Locktime             string `json:"locktime"`
	Deletedate           string `json:"deletedate"`
	Isadmin              int    `json:"isadmin"`
	Createdate           string `json:"createdate"`
	Createuserid         string `json:"createuserid"`
	Modifydate           string `json:"modifydate"`
	Modifyuserid         string `json:"modifyuserid"`
	ZDesc                string `json:"zDesc"`
	Islinlay             int    `json:"islinlay"`
	Pwderrorunm          int32  `json:"pwderrorunm"`
	Pwderrortime         string `json:"pwderrortime"`
	Modifypwddate        string `json:"modifypwddate"`
	Roleid               int    `json:"roleid"`
	Expiretime           string `json:"expiretime"`
	Ruleid               int    `json:"ruleid"`
	DicCityId            int    `json:"dicCityId"`
	PostId               int    `json:"postId"`
	DepartmentId         int    `json:"departmentId"`
	EmployeeNumber       string `json:"employeeNumber"`
	DomainPassword       string `json:"domainPassword"`
	UserPasswordPolicyId int    `json:"userPasswordPolicyId"`
	UserAuthenPolicyId   int    `json:"userAuthenPolicyId"`
	AuthKey              string `json:"authKey"`
	QrCode               string `json:"qrCode"`
}
type UserAll struct {
	Uid                  string `gorm:"primary_key" json:"uid"`
	Uname                string `json:"uname"`
	Sex                  string `json:"sex"`
	Cardid               string `json:"cardid"`
	Birthdate            string `json:"birthdate"`
	Mail                 string `json:"mail"`
	Mobile               string `json:"mobile"`
	Password             string `json:"password"`
	Status               string `json:"status"`
	Lastlogin            string `json:"lastlogin"`
	Locktime             string `json:"locktime"`
	Deletedate           string `json:"deletedate"`
	Isadmin              int    `json:"isadmin"`
	Createdate           string `json:"createdate"`
	Createuserid         string `json:"createuserid"`
	Modifydate           string `json:"modifydate"`
	Modifyuserid         string `json:"modifyuserid"`
	ZDesc                string `json:"zDesc"`
	Islinlay             int    `json:"islinlay"`
	Pwderrorunm          int32  `json:"pwderrorunm"`
	Pwderrortime         string `json:"pwderrortime"`
	Modifypwddate        string `json:"modifypwddate"`
	Roleid               int    `json:"roleid"`
	Expiretime           string `json:"expiretime"`
	Ruleid               int    `json:"ruleid"`
	DicCityId            int    `json:"dicCityId"`
	PostId               int    `json:"postId"`
	DepartmentId         int    `json:"departmentId"`
	EmployeeNumber       string `json:"employeeNumber"`
	DomainPassword       string `json:"domainPassword"`
	UserPasswordPolicyId int    `json:"userPasswordPolicyId"`
	UserAuthenPolicyId   int    `json:"userAuthenPolicyId"`
	AuthKey              string `json:"authKey"`
	Rolename             string `json:"rolename"`
	DepName              string `json:"depName"`
	QrCode               string `json:"qrCode"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Createdate = time.Now().Format("2006-01-02 15:04:05")
	user.Lastlogin = time.Now().Format("0000-00-00 00:00:00")
	user.Locktime = time.Now().Format("0000-00-00 00:00:00")
	user.Deletedate = time.Now().Format("0000-00-00 00:00:00")
	user.Modifydate = time.Now().Format("0000-00-00 00:00:00")
	user.Pwderrortime = time.Now().Format("0000-00-00 00:00:00")
	user.Modifypwddate = time.Now().Format("0000-00-00 00:00:00")
	return
}
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	user.Modifydate = time.Now().Format("2006-01-02 15:04:05")
	return
}
func (user *User) BeforeDel(tx *gorm.DB) (err error) {
	user.Deletedate = time.Now().Format("2006-01-02 15:04:05")
	return
}

type Expdp struct {
	Uid    string `gorm:"primary_key" json:"uid"`
	Uname  string `json:"uname"`
	Roleid int    `json:"roleid"`
	//Sex         	 		 string    `json:"sex"`
	//Birthdate   	  		 string    `json:"birthdate"`
	//Cardid      	 		 string    `json:"cardid"`
	Mobile         string `json:"mobile"`
	DicCityId      int    `json:"dic_city_id"`
	DepartmentId   int    `json:"department_id"`
	EmployeeNumber string `json:"employee_number"`
	Mail           string `json:"mail"`
}

type ExpdpAll struct {
	Uid   string `gorm:"primary_key" json:"uid"`
	Uname string `json:"uname"`
	Role  string `json:"roleid"`
	//Sex         	 		 string    `json:"sex"`
	//Birthdate   	  		 string    `json:"birthdate"`
	//Cardid      	 		 string    `json:"cardid"`
	Mobile         string `json:"mobile"`
	DicCity        string `json:"dic_city_id"`
	Department     string `json:"department_id"`
	EmployeeNumber string `json:"employee_number"`
	Mail           string `json:"mail"`
}

func (User) TableName() string {
	return "z_user"
}

func (UserAll) TableName() string {
	return "z_user"
}

func (Expdp) TableName() string {
	return "z_user"
}
