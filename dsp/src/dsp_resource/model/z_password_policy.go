package model

import (
	"gorm.io/gorm"
	"time"
)

/**
* @author tianqiang
* @date 2021/7/19.
 */
type PasswordPolicy struct {
	Id                   int    `gorm:"primary_key" json:"id"`
	PolicyName           string `json:"policyName"`
	PolicyType           string `json:"policyType"`
	PolicyLengthMore     int    `json:"policyLengthMore"`
	PolicyLengthLess     int    `json:"policyLengthLess"`
	PolicyIncludeChar    string `json:"policyIncludeChar"`
	CreateDate           string `json:"createDate"`
	CreateUserId         string `json:"createUserId"`
	ModifyDate           string `json:"modifyDate"`
	ModifyUserId         string `json:"modifyUserId"`
	PasswordErrorNumLock int    `json:"passwordErrorNumLock"`
	PolicyValidity       int    `json:"policyValidity"`
	PasswordLength       int    `json:"passwordLength"`
	LockDuration         int    `json:"lockDuration"`
}

func (PasswordPolicy) TableName() string {
	return "z_password_policy"
}

func (policy *PasswordPolicy) BeforeCreate(tx *gorm.DB) (err error) {
	policy.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	policy.ModifyDate = time.Now().Format("0000-00-00 00:00:00")
	return
}

func (policy *PasswordPolicy) BeforeSave(tx *gorm.DB) (err error) {
	policy.ModifyDate = time.Now().Format("2006-01-02 15:04:05")
	return
}
