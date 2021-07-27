package model

import "time"

/**
* @author tianqiang
* @date 2021/7/19.
 */
type Store struct {
	ResId                     int    `gorm:"primary_key" json:"resId"`
	StoreName                 string `json:"storeName"`
	StoreType                 int    `json:"storeType"`
	StoreVersion              int    `json:"storeVersion"`
	StorePort                 int    `json:"storePort"`
	ResHostId                 int    `json:"resHostId"`
	StoreClass                string `json:"storeClass"`
	StoreUrl                  string `json:"storeUrl"`
	StoreSid                  string `json:"storeSid"`
	StoreAdminAccount         string `json:"storeAdminAccount"`
	StoreAdminAccountPassword string `json:"storeAdminAccountPassword"`
}

type StoreParams struct {
	ResName      string `json:"resName"`
	ResIpV4      string `json:"resIpV4"`
	ResIpV6      string `json:"resIpV6"`
	DepartmentId int    `json:"departmentId"`
	StoreType    int    `json:"storeType"`
}

type StoreAll struct {
	Id                        int       `gorm:"primary_key" json:"id"`
	ResName                   string    `json:"resName"`
	ResIpV4                   string    `json:"resIpV4"`
	ResIpV6                   string    `json:"resIpV6"`
	DicConnectId              int       `json:"dicConnectId"`
	DicTypeId                 int       `json:"dicTypeId"`
	DicCityId                 int       `json:"dicCityId"`
	DicManufactorId           int       `json:"dicManufactorId"`
	DepartmentId              int       `json:"departmentId"`
	ResManagerUserId          string    `json:"resManagerUserId"`
	Status                    int       `json:"status"`
	ZDesc                     string    `json:"zDesc"`
	CreateDate                time.Time `json:"createDate"`
	CreateUserId              string    `json:"createUserId"`
	ModifyDate                time.Time `json:"modifyDate"`
	ModifyUserId              string    `json:"modifyUserId"`
	ResPasswordPolicyId       int       `json:"resPasswordPolicyId"`
	SecurePort                int       `json:"securePort"`
	SomfIp                    string    `json:"somfIp"`
	ResId                     int       `gorm:"primary_key" json:"resId"`
	StoreName                 string    `json:"storeName"`
	StoreType                 int       `json:"storeType"`
	StoreVersion              int       `json:"storeVersion"`
	StorePort                 int       `json:"storePort"`
	ResHostId                 int       `json:"resHostId"`
	StoreClass                string    `json:"storeClass"`
	StoreUrl                  string    `json:"storeUrl"`
	StoreSid                  string    `json:"storeSid"`
	StoreAdminAccount         string    `json:"storeAdminAccount"`
	StoreAdminAccountPassword string    `json:"storeAdminAccountPassword"`
}

type StoreInfo struct {
	Res   `mapstructure:",squash"`
	Store `mapstructure:",squash"`
}

func (Store) TableName() string {
	return "z_res_store"
}

//导出存储设备的封装类
type ExportStore struct {
	ResName           string `json:"resName"`           //资源名称 res
	ResIpV4           string `json:"resIpV4"`           //ipv4 res
	ResIpV6           string `json:"resIpV6"`           //ipv6 res
	ResTypeName       string `json:"resTypeName"`       //资源类型 res 字典
	DepName           string `json:"depName"`           //部门 z_department
	ResAccount        string `json:"resAccount"`        //资源账号  account
	ResAccountName    string `json:"resAccountName"`    //资源账号名称  account
	ResPasswordPolicy string `json:"resPasswordPolicy"` //密码策略 account 关联密码策略表
	DicConnectId      string `json:"dicConnectId"`      //是否可连接 res 是否
	ManufactorName    string `json:"manufactorName"`    //生产厂家 res 字典
	StoreName         string `json:"storeName"`         //数据库名
	ResHostName       string `json:"resHostName"`       //物理主机 res_host_id
	StoreType         string `json:"storeType"`         //存储类型
	StoreVersion      string `json:"storeVersion"`      //存储版本
	StoreSid          string `json:"storeSid"`          //存储实例
	StoreUrl          string `json:"storeUrl"`          //存储设备连接url
	StorePort         string `json:"storePort"`         //存储端口
	StoreClass        string `json:"storeClass"`        //存储设备驱动类名
	DicCity           string `json:"dicCity"`           //城市 z_dictionary

}

type StorePage struct {
	Res           `mapstructure:",squash"`
	Store         `mapstructure:",squash"`
	StoreTypeName string `json:"storeTypeName"`
	DepName       string `json:"depName"`
	PolicyName    string `json:"policyName"`
	HostName      string `json:"hostName"`
	Uname         string `json:"uname"`
}
