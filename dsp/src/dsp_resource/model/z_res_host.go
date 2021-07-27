package model

import (
	"time"
)

/**
* @author tianqiang
* @date 2021/7/19.
 */
type Host struct {
	ResId             int    `gorm:"primary_key" json:"resId"`
	HostType          int    `json:"hostType"`
	HostSystemType    int    `json:"hostSystemType"`
	HostSystemVersion int    `json:"hostSystemVersion"`
	ProtocolType      int    `json:"protocolType"`
	ProtocolPort      int    `json:"protocolPort"`
	IsSudo            int    `json:"isSudo"`
	AdminAccount      string `json:"adminAccount"`
	AdminPassword     string `json:"adminPassword"`
	AdminConSym       string `json:"adminConSym"`
	UnixSuType        string `json:"unix_su_type"`
	UnixRootPassword  string `json:"unixRootPassword"`
	UnixRootConSym    string `json:"unixRootConSym"`
	UnixResSkipId     int    `json:"unixResSkipId"`
	WinType           int    `json:"winType"`
	WinDomainName     string `json:"winDomainName"`
	WinDomainDn       string `json:"winDomainDn"`
	WinDomainResId    int    `json:"winDomainResId"`
	DicVersionId      int    `json:"dicVersionId"`
	DicModelId        int    `json:"dicModelId"`
	RdpPort           int    `json:"rdpPort"`
}

type HostParams struct {
	ResName      string `json:"resName"`
	ResIpV4      string `json:"resIpV4"`
	ResIpV6      string `json:"resIpV6"`
	HostType     int    `json:"hostType"`
	DepartmentId int    `json:"departmentId"`
}

type HostAll struct {
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
	ResId               int       `json:"resId"`
	HostType            int       `json:"hostType"`
	HostSystemType      int       `json:"hostSystemType"`
	HostSystemVersion   int       `json:"hostSystemVersion"`
	ProtocolType        int       `json:"protocolType"`
	ProtocolPost        int       `json:"protocolPost"`
	IsSudo              int       `json:"isSudo"`
	AdminAccount        string    `json:"adminAccount"`
	AdminPassword       string    `json:"adminPassword"`
	AdminConSym         string    `json:"adminConSym"`
	UnixRootPassword    string    `json:"unixRootPassword"`
	UnixRootConSym      string    `json:"unixRootConSym"`
	UnixResSkipId       int       `json:"unixResSkipId"`
	WinType             int       `json:"winType"`
	WinDomainName       string    `json:"winDomainName"`
	WinDomainDn         string    `json:"winDomainDn"`
	WinDomainResId      int       `json:"winDomainResId"`
	DicVersionId        int       `json:"dicVersionId"`
	DicModelId          int       `json:"dicModelId"`
	RdpPort             int       `json:"rdpPort"`
}

type HostInfo struct {
	Res  `mapstructure:",squash"`
	Host `mapstructure:",squash"`
}
type HostLogs struct {
	ResName  string `json:"resName"`
	ResIpV4  string `json:"resIpV4"`
	ErrorLog string `json:"errorLog"`
}

func (Host) TableName() string {
	return "z_res_host"
}

//导出主机设备的封装类
type ExportHost struct {
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
	HostType          string `json:"hostType"`          //主机类型 z_dictionary
	WinType           string `json:"winType"`           //win主机类型
	HostSystemType    string `json:"hostSystemType"`    //主机系统类型
	HostSystemVersion string `json:"hostSystemVersion"` //主机系统版本
	ProtocolType      string `json:"protocolType"`      //协议类型
	ProtocolPort      string `json:"protocolPort"`      //协议端口
	UnixResSkipName   string `json:"unixResSkipName"`   //跳转机的名称
	WinDomainName     string `json:"winDomainName"`     //域名
	WinDomainDn       string `json:"winDomainDn"`       //域控主机的dn
	WinDomainResName  string `json:"winDomainResName"`  //域主机名称
	DicVersionId      string `json:"dicVersionId"`      //版本
	DicModelId        string `json:"dicModelId"`        //型号
	DicCity           string `json:"dicCity"`           //城市 z_dictionary

}

type HostPage struct {
	Res          `mapstructure:",squash"`
	Host         `mapstructure:",squash"`
	HostTypeName string `json:"hostTypeName"`
	DepName      string `json:"depName"`
	PolicyName   string `json:"policyName"`
	SkipName     string `json:"skipName"`
	Uname        string `json:"uname"`
	DomainName   string `json:"domainName"`
}
