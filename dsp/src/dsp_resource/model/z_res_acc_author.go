package model

import (
	"gorm.io/gorm"
	"time"
)

/**
* @author tianqiang
* @date 2021/7/19.
 */
type AccountAuthor struct {
	Id             int64  `gorm:"primary_key" json:"id"`
	ResAccountId   int64  `json:"resAccountId"`
	UserId         string `json:"userId"`
	ResId          int64  `json:"resId"`
	LoginType      int64  `json:"loginType"`
	OrderCollectId int64  `json:"orderCollectId"`
	BeginDate      string `json:"beginDate"`
	EndDate        string `json:"endDate"`
	ZDesc          string `json:"zDesc"`
	AccessCount    int64  `json:"accessCount"`
	CreateDate     string `json:"createDate"`
}

func (AccountAuthor) TableName() string {
	return "z_res_acc_author"
}

type ExportAccountAuthor struct {
	UserName         string `json:"userName"`         //被授权人 z_user user_id
	BeginDate        string `json:"beginDate"`        //开始时间
	EndDate          string `json:"endDate"`          //结束时间
	LoginType        string `json:"loginType"`        //登录类型
	OrderCollectName string `json:"orderCollectName"` //命令集
	ClientName       string `json:"clientName"`       //客户端列表
	ZDesc            string `json:"zDesc"`            //描述
}

func (account *AccountAuthor) BeforeCreate(tx *gorm.DB) (err error) {
	account.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	return
}

type AccountAuthorPage struct {
	AccountAuthor `mapstructure:",squash"`
	UserName      string `json:"userName"`
	ClientId      string `json:"clientId"`
}
