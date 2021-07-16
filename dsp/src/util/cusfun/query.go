package cusfun

import (
	"gorm.io/gorm"
)


// 查询信息
func QueryBy(db *gorm.DB,id string, model interface{}) (rst interface{},err error) {
	err = db.Find(&model,"id = ?",id).Error
	return rst,err
}


