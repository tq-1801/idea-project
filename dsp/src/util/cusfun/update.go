package cusfun

import (
	"gorm.io/gorm"
)

// 修改多个字段
func Update(db *gorm.DB, model interface{}) (rst interface{},err error){
	err = db.Save(&model).Error
	return rst,err
}


