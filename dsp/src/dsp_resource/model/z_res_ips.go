package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type ResIps struct {
	Id    int64  `gorm:"primary_key"  json:"id"`
	ResId int64  `json:"resId"`
	IpV4  string `json:"ipV4"`
	IpV6  string `json:"ipV6"`
}

func (ResIps) TableName() string {
	return "z_res_ips"
}
