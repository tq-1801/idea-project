package model

import (
	"gorm.io/gorm"
	"time"
)

/**
 * @Author: tianqiang
 * @Date: 2021/7/15
 */
type AuthenticactionPolicy struct {
	Id                        int64  `gorm:"primary_key" json:"id"`
	AuthenticactionPolicyName string `json:"authenticactionPolicyName"`
	LoginBeginTime            string `json:"loginBeginTime"`
	LoginEndTime              string `json:"loginEndTime"`
	LoginIp                   string `json:"loginIp"`
	CreateDate                string `json:"createDate"`
	ModifyDate                string `json:"modifyDate"`
	CreateUserId              string `json:"createUserId"`
	ModifyUserId              string `json:"modifyUserId"`
	ZDesc                     string `json:"zdesc"`
	ResTerminalId             string `json:"resTerminalId"`
	AuthenType                int64  `json:"authenType"`
	LoginBeginDate            string `json:"loginBeginDate"`
	LoginEndDate              string `json:"loginEndDate"`
	IsInlay                   int    `json:"isInlay"`
}

func (AuthenticactionPolicy) TableName() string {
	return "z_authenticaction_policy"
}

func (policy *AuthenticactionPolicy) BeforeCreate(tx *gorm.DB) (err error) {
	policy.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	policy.ModifyDate = time.Now().Format("0000-00-00 00:00:00")
	return
}

func (policy *AuthenticactionPolicy) BeforeSave(tx *gorm.DB) (err error) {
	policy.ModifyDate = time.Now().Format("2006-01-02 15:04:05")
	return
}
