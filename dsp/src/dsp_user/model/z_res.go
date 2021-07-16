package model

import (
	"time"
)

/**

* @author tianqiang

* @date 2021/3/22.

 */

type Res struct {
	Id                  int       `gorm:"primary_key" json:"id"`
	ResName             string    `json:"resName"`
	ResIpV4             string    `json:"resIpV4"`
	ResIpV6             string    `json:"resIpV6"`
	DicConnectId        int       `json:"dicConnectId"`
	DicTypeId           int       `json:"dicTypeId"`
	DicCityId           int       `json:"dicCityId"`
	DicManufactorId     int       `json:"dicManufactorId"`
	DepartmentId        int       `json:"departmentId"`
	ResManagerUserId    string    `json:"resManagerUserId"`
	Status              int       `json:"status"`
	ZDesc               string    `json:"zDesc"`
	CreateDate          time.Time `json:"createDate"`
	CreateUserId        string    `json:"createUserId"`
	ModifyDate          time.Time `json:"modifyDate"`
	ModifyUserId        string    `json:"modifyUserId"`
	ResPasswordPolicyId int       `json:"resPasswordPolicyId"`
	SecurePort          int       `json:"securePort"`
	SomfIp              string    `json:"somfIp"`
}

func (Res) TableName() string {
	return "z_res"
}
