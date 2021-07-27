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
	Uid                  string `gorm:"primary_key" json:"uid" mapstructure:"uid" validate:"" trans:"用户名"`
	Uname                string `json:"uname" mapstructure:"uname" validate:"require" trans:"姓名"`
	Sex                  string `json:"sex" mapstructure:"sex" validate:"" trans:"性别"`
	Cardid               string `json:"cardid" mapstructure:"cardid" validate:"" trans:"身份证号码"`
	Birthdate            string `json:"birthdate" mapstructure:"birthdate" validate:"" trans:"出生日期"`
	Mail                 string `json:"mail" mapstructure:"mail" validate:"email" trans:"邮箱"`
	Mobile               string `json:"mobile" mapstructure:"mobile" validate:"" trans:"手机号码"`
	Password             string `json:"password" mapstructure:"password" validate:"" trans:"人员登录密码"`
	Status               string `json:"status" mapstructure:"status" validate:"" trans:"用户状态"`
	Lastlogin            string `json:"lastlogin" mapstructure:"lastlogin" validate:"" trans:"用户最后登录时间"`
	Locktime             string `json:"locktime" mapstructure:"locktime" validate:"" trans:"用户锁定时间"`
	Deletedate           string `json:"deletedate" mapstructure:"deletedate" validate:"" trans:"用户删除时间"`
	Isadmin              int    `json:"isadmin" mapstructure:"isadmin" validate:"" trans:"用户权限"`
	Createdate           string `json:"createdate" mapstructure:"createdate" validate:"" trans:"数据创建时间"`
	Createuserid         string `json:"createuserid" mapstructure:"createuserid" validate:"" trans:"数据创建用户账号"`
	Modifydate           string `json:"modifydate" mapstructure:"modifydate" validate:"" trans:"数据修改时间"`
	Modifyuserid         string `json:"modifyuserid" mapstructure:"modifyuserid" validate:"" trans:"数据修改用户账号"`
	ZDesc                string `json:"zDesc" mapstructure:"zDesc" validate:"" trans:"描述"`
	Islinlay             int    `json:"islinlay" mapstructure:"islinlay" validate:"" trans:"是否初始化数据"`
	Pwderrorunm          int32  `json:"pwderrorunm" mapstructure:"pwderrorunm" validate:"" trans:"密码错误次数"`
	Pwderrortime         string `json:"pwderrortime" mapstructure:"pwderrortime" validate:"" trans:"密码错误时间"`
	Modifypwddate        string `json:"modifypwddate" mapstructure:"modifypwddate" validate:"" trans:"密码修改时间"`
	Roleid               int    `json:"roleid" mapstructure:"roleid" validate:"" trans:"角色id"`
	Expiretime           string `json:"expiretime" mapstructure:"expiretime" validate:"" trans:"过期日期"`
	Ruleid               int    `json:"ruleid" mapstructure:"ruleid" validate:"" trans:"规则id"`
	DicCityId            int    `json:"dicCityId" mapstructure:"dicCityId" validate:"" trans:"城市字典"`
	PostId               int    `json:"postId" mapstructure:"postId" validate:"" trans:"岗位"`
	DepartmentId         int    `json:"departmentId" mapstructure:"departmentId" validate:"require" trans:"部门"`
	EmployeeNumber       string `json:"employeeNumber" mapstructure:"employeeNumber" validate:"" trans:"员工编号"`
	DomainPassword       string `json:"domainPassword" mapstructure:"domainPassword" validate:"" trans:"用户在域控上的密码"`
	UserPasswordPolicyId int    `json:"userPasswordPolicyId" mapstructure:"userPasswordPolicyId" validate:"" trans:"用户密码策略ID"`
	UserAuthenPolicyId   int    `json:"userAuthenPolicyId" mapstructure:"userAuthenPolicyId" validate:"" trans:"认证策略id"`
	AuthKey              string `json:"authKey" mapstructure:"authKey" validate:"" trans:"用户软令牌加密密钥"`
	QrCode               string `json:"qrCode" mapstructure:"qrCode" validate:"" trans:"二维码"`
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
