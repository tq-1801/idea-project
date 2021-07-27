package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type HostProtocol struct {
	Id           int `gorm:"primary_key" json:"id"`
	ResId        int `json:"resId"`
	ProtocolType int `json:"protocolType"`
	ProtocolPort int `json:"protocolPort"`
}

func (HostProtocol) TableName() string {
	return "z_res_host_protocol"
}

type HostProtocolPage struct {
	Id           int    `gorm:"primary_key" json:"id"`
	ResId        int    `json:"resId"`
	ProtocolType int    `json:"protocolType"`
	ProtocolPort int    `json:"protocolPort"`
	DicName      string `json:"dicName"`
}
