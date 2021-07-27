package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type AccAuthClient struct {
	Id        int64 `gorm:"primary_key" json:"id"`
	ResAuthId int64 `json:"resAuthId"`
	ClientId  int64 `json:"clientId"`
}

func (AccAuthClient) TableName() string {
	return "z_res_acc_auth_client"
}
