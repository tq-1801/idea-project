package model

type Example struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Memo string `json:"memo"`
}

func (Example) TableName() string {
	return "z_example"
}
