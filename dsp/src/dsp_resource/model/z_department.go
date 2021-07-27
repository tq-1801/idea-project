package model

/**
* @author tianqiang
* @date 2021/7/19.
 */
type Department struct {
	Id               int    `gorm:"AUTO_INCREMENT" json:"id"`
	DepName          string `json:"depName"`
	DepNumber        string `json:"depNumber"`
	DicCityId        int    `json:"dicCityId"`
	ManagerUserId    string `json:"managerUserId"`
	IsLeaf           int    `json:"isLeaf"`
	ZDesc            string `json:"zDesc"`
	IsInlay          int    `json:"isInlay"`
	CreateDate       string `json:"createDate"`
	CreateUserId     string `json:"createUserId"`
	ModifyDate       string `json:"modifyDate"`
	ModifyUserId     string `json:"modifyUserId"`
	SupId            int    `json:"supId"`
	TreePath         string `json:"treePath"`
	PasswordPolicyId int    `json:"passwordPolicyId"`
	TerminalPolicyId int    `json:"terminalPolicyId"`
}

func (Department) TableName() string {
	return "z_department"
}
