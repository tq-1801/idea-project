package model

/**

* @author tianqiang

* @date 2021/4/10.

 */
type PasswordPolicy struct {
	Id                   int    `gorm:"primary_key" json:"id"`
	PolicyName           string `json:"policyName"`
	PolicyType           string `json:"policyType"`
	PolicyLengthMore     int64  `json:"policyLengthMore"`
	PolicyLengthLess     int64  `json:"policyLengthLess"`
	PolicyIncludeChar    string `json:"policyIncludeChar"`
	CreateDate           string `json:"createDate"`
	CreateUserId         string `json:"createUserId"`
	ModifyDate           string `json:"modifyDate"`
	ModifyUserId         string `json:"modifyUserId"`
	PasswordErrorNumLock int64  `json:"passwordErrorNumLock"`
	PolicyValidity       int    `json:"policyValidity"`
	PasswordLength       int64  `json:"passwordLength"`
	LockDuration         int    `json:"lockDuration"`
}

func (PasswordPolicy) TableName() string {
	return "z_password_policy"
}

//type UserPwd struct {
//	Uid         		     string    `gorm:"primary_key" json:"uid"`
//	Uname        			 string    `json:"uname"`
//	Password     			 string    `json:"password"`
//	Expiredate   			 string       `json:"expiredate"`
//	UserPasswordPolicyId     int       `json:"userPasswordPolicyId"`
//	PolicyValidity 			int		`json:"policyValidity"`
//}
//
//func (UserPwd) TableName() string {
//	return "z_password_policy"
//}
