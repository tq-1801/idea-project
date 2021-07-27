package model

import (
	"gorm.io/gorm"
	"time"
)

/**
* @author tianqiang
* @date 2021/7/19.
 */
type Res struct {
	Id                  int    `gorm:"primary_key" json:"id"`
	ResName             string `json:"resName"`
	ResIpV4             string `json:"resIpV4"`
	ResIpV6             string `json:"resIpV6"`
	DicConnectId        int    `json:"dicConnectId"`
	DicTypeId           int    `json:"dicTypeId"`
	DicCityId           int    `json:"dicCityId"`
	DicManufactorId     int    `json:"dicManufactorId"`
	DepartmentId        int    `json:"departmentId"`
	ResManagerUserId    string `json:"resManagerUserId"`
	Status              int    `json:"status"`
	ZDesc               string `json:"zDesc"`
	CreateDate          string `json:"createDate"`
	CreateUserId        string `json:"createUserId"`
	ModifyDate          string `json:"modifyDate"`
	ModifyUserId        string `json:"modifyUserId"`
	ResPasswordPolicyId int    `json:"resPasswordPolicyId"`
	SecurePort          int    `json:"securePort"`
	SomfIp              string `json:"somfIp"`
}

func (Res) TableName() string {
	return "z_res"
}

func (res *Res) BeforeCreate(tx *gorm.DB) (err error) {
	res.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	res.ModifyDate = time.Now().Format("0000-00-00 00:00:00")
	return
}

func (res *Res) BeforeSave(tx *gorm.DB) (err error) {
	res.ModifyDate = time.Now().Format("2006-01-02 15:04:05")
	return
}

/**
 * @Description:导入错误信息封装类
 */
type ImportErrorLogs struct {
	Name          string `json:"name"`
	SecondMessage string `json:"secondMessage"`
	ErrorLog      string `json:"errorLog"`
}

type ICMP struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}
