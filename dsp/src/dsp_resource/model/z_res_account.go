package model

import (
	"gorm.io/gorm"
	"time"
)
/**
* @author tianqiang
* @date 2021/7/19.
 */

/**
绑定账号信息
*/
type Account struct {
	Id                 int    `gorm:"primary_key" json:"id"`
	ResId              int    `json:"resId"`
	ResAccount         string `json:"resAccount"`
	ResAccountName     string `json:"resAccountName"`
	ResAccountPassword string `json:"resAccountPassword"`
	ResAccountStatus   int    `json:"resAccountStatus"`
	ResAccountYum      string `json:"resAccountYum"`
	ModifyPasswordDate string `json:"modifyPasswordDate"`
	CreateDate         string `json:"createDate"`
	CreateUserId       string `json:"createUserId"`
	ModifyDate         string `json:"modifyDate"`
	ModifyUserId       string `json:"modifyUserId"`
	ZDesc              string `json:"zDesc"`
	ResAccountType     int    `json:"resAccountType"`
	SuType             string `json:"suType"`
	IsAdmin            int    `json:"isAdmin"`
	IsSuper            int    `json:"isSuper"`
	IsAuto             int    `json:"isAuto"`
}

func (Account) TableName() string {
	return "z_res_account"
}

type ExportAccount struct {
	ResAccount         string `json:"resAccount"`         //账号
	ResAccountName     string `json:"resAccountName"`     //账号名称
	ResAccountStatus   string `json:"resAccountStatus"`   //账号状态
	ResAccountYum      string `json:"resAccountYum"`      //登录回显
	CreateDate         string `json:"createDate"`         //创建时间
	ResAccountTypeName string `json:"resAccountTypeName"` //账号类型
	ZDesc              string `json:"zDesc"`              //描述
	SuType             string `json:"suType"`             //su类型
	IsAdmin            string `json:"isAdmin"`            //是否管理员

}

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	account.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	account.ModifyDate = time.Now().Format("0000-00-00 00:00:00")
	return
}

func (account *Account) BeforeSave(tx *gorm.DB) (err error) {
	account.ModifyDate = time.Now().Format("2006-01-02 15:04:05")
	return
}

type AccountPage struct {
	Account       `mapstructure:",squash"`
	AccountStatus string `json:"resAccountStatusName"`
	IsAdminName   string `json:"isAdminName"`
}

type AccountInfo struct {
	ResIpV4            string `json:"resIpV4"`
	ResAccount         string `json:"resAccount"`
	ResAccountPassword string `json:"resAccountPassword"`
	ResAccountId       int    `json:"resAccountId"`
	ResPort            string `json:"resPort"`
	IsAdmin            int    `json:"isAdmin"`
}
