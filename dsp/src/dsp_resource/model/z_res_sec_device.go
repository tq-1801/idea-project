package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type ResSecDeviceInfo struct {
	//主表信息
	Res `mapstructure:",squash"`
	//从表信息
	ResSecDevice `mapstructure:",squash"`
}

//导出安全设备的封装类
type ExportSecDevice struct {
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
	SecDeviceType     string `json:"secDeviceType"`     //安全设备类型，字典
	ResSkip           string `json:"resSkip"`           //跳转机名称,待确认和哪个表进行关联
	ProtocolType      string `json:"protocolType"`      //协议类型，字典
	ProtocolPort      int64  `json:"protocolPort"`      //协议端口
	DicVersion        string `json:"dicVersion"`        //版本
	DicModel          string `json:"dicModel"`          //型号
	DicCity           string `json:"dicCity"`           //城市 字典

}

//安全设备表
type ResSecDevice struct {
	ResId         int64  `gorm:"primary_key" json:"resId"`
	SecDeviceType int64  `json:"secDeviceType"`
	AdminAccount  string `json:"adminAccount"`
	AdminPassword string `json:"adminPassword"`
	RootPassword  string `json:"rootPassword"`
	AdminConSym   string `json:"adminConSym"`
	RootConSym    string `json:"rootConSym"`
	ResSkipId     int64  `json:"resSkipId"`
	WebUrl        string `json:"webUrl"`
	ProtocolType  int64  `json:"protocolType"`
	ProtocolPort  int64  `json:"protocolPort"`
	DicVersionId  int64  `json:"dicVersionId"`
	DicModelId    int64  `json:"dicModelId"`
}

func (ResSecDevice) TableName() string {
	return "z_res_sec_device"
}

//业务系统信息表
type ResSecDevicePage struct {
	//主表信息
	Res `mapstructure:",squash"`
	//从表信息
	ResSecDevice `mapstructure:",squash"`
	//关联中文信息
	SecDeviceTypeName string `json:"secDeviceTypeName"`
	DepName           string `json:"depName"`
	PolicyName        string `json:"policyName"`
	SkipName          string `json:"skipName"`
	Uname             string `json:"uname"`
}
