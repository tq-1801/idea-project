package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type ResNetworkDevice struct {
	ResId             int    `gorm:"primary_key" json:"resId"`
	NetworkDeviceType int    `json:"networkDeviceType"`
	AdminAccount      string `json:"adminAccount"`
	AdminPassword     string `json:"adminPassword"`
	RootPassword      string `json:"rootPassword"`
	AdminConSym       string `json:"adminConSym"`
	RootConSym        string `json:"rootConSym"`
	IsReturn          int    `json:"isReturn"`
	IsKvm             int    `json:"isKvm"`
	ResSkipId         int    `json:"resSkipId"`
	ProtocolType      int    `json:"protocolType"`
	ProtocolPort      int    `json:"protocolPort"`
	DicVersionId      int    `json:"dicVersionId"`
	DicModelId        int    `json:"dicModelId"`
}

type ResNetworkDeviceInfo struct {
	Res              `mapstructure:",squash"`
	ResNetworkDevice `mapstructure:",squash"`
}

func (ResNetworkDevice) TableName() string {
	return "z_res_network_device"
}

//导出存储设备的封装类
type ExportNetwork struct {
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
	NetworkDeviceType string `json:"networkDeviceType"` //网络设备类型
	IsReturn          string `json:"isReturn"`          //是否回车
	IsKvm             string `json:"isKvm"`             //是否kvm
	ResSkipName       string `json:"resSkipName"`       //跳转机名称
	ProtocolType      string `json:"protocolType"`      //协议类型
	ProtocolPort      string `json:"protocolPort"`      //协议端口
	DicCity           string `json:"dicCity"`           //城市 z_dictionary

}

type ResNetworkDevicePage struct {
	Res                   `mapstructure:",squash"`
	ResNetworkDevice      `mapstructure:",squash"`
	NetworkDeviceTypeName string `json:"networkDeviceTypeName"`
	DepName               string `json:"depName"`
	PolicyName            string `json:"policyName"`
	SkipName              string `json:"skipName"`
	Uname                 string `json:"uname"`
}
