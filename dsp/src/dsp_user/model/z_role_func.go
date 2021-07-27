package model

type UgFunc struct {
	Id        int `gorm:"primary_key"`
	Roleid    int
	Funcid    int
	Level     int
	Oper      string
	FuncSupid int
}

func (UgFunc) TableName() string {
	return "z_role_func"
}
