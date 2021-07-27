package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type ResBusinessInfo struct {
	//主表信息
	Res `mapstructure:",squash"`
	//从表信息
	ResBusiness `mapstructure:",squash"`
}

//导出业务系统的封装类
type ExportBusiness struct {
	ResName              string `json:"resName"`              //资源名称 res
	ResIpV4              string `json:"resIpV4"`              //ipv4 res
	ResIpV6              string `json:"resIpV6"`              //ipv6 res
	ResTypeName          string `json:"resTypeName"`          //资源类型 res 字典
	DepName              string `json:"depName"`              //部门 z_department
	ResAccount           string `json:"resAccount"`           //资源账号  account
	ResAccountName       string `json:"resAccountName"`       //资源账号名称  account
	ResPasswordPolicy    string `json:"resPasswordPolicy"`    //密码策略 account 关联密码策略表
	DicConnectId         string `json:"dicConnectId"`         //是否可连接 res 是否
	ManufactorName       string `json:"manufactorName"`       //生产厂家 res 字典
	Url                  string `json:"url"`                  //业务系统URL
	ResHostName          string `json:"resHostName"`          //物理主机名称 z_res_host
	ConnectStoreName     string `json:"connectStoreName"`     //连接数据库名称 z_res_store
	ResBusinessType      string `json:"resBusinessType"`      //业务系统类型 z_dictionary
	SupBusinessName      string `json:"supBusinessName"`      //父应用名称 z_res_business
	BusinessRelationName string `json:"businessRelationName"` //业务名称关联类型 z_dictionary
	LoginTypeName        string `json:"loginTypeName"`        //登录类型 z_dictionary
	ConnectTypeName      string `json:"connectTypeName"`      //连接类型 z_dictionary
	LoginProtocol        string `json:"loginProtocol"`        //登录协议 z_dictionary
	BsAccount            string `json:"bsAccount"`            //账号属性名
	DicCity              string `json:"dicCity"`              //城市 z_dictionary

}

//业务系统信息表
type ResBusiness struct {
	ResId                int64  `gorm:"primary_key" json:"resId"`
	ResHostId            int64  `json:"resHostId"`
	Url                  string `json:"url"`
	ResBusinessType      int64  `json:"resBusinessType"`
	FrontMess            string `json:"frontMess"`
	LastMess             string `json:"lastMess"`
	BusinessRelationType int64  `json:"businessRelationType"`
	LoginType            int64  `json:"loginType"`
	ConnectType          int64  `json:"connectType"`
	ConnectStoreId       int64  `json:"connectStoreId"`
	BsAccount            string `json:"bsAccount"`
	BsPassword           string `json:"bsPassword"`
	BsButton             string `json:"bsButton"`
	LoginProtocol        int64  `json:"loginProtocol"`
	SupBusinessId        int64  `json:"supBusinessId"`
}

func (ResBusiness) TableName() string {
	return "z_res_business"
}

//业务系统信息表
type ResBusinessPage struct {
	//主表信息
	Res `mapstructure:",squash"`
	//从表信息
	ResBusiness `mapstructure:",squash"`
	//关联中文信息
	ResBusinessTypeName string `json:"resBusinessTypeName"`
}
