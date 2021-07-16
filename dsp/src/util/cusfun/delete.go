package cusfun

import (
	"gorm.io/gorm"
)

// 删除
func Delete(db *gorm.DB, model interface{}) (rst interface{},err error) {
	err = db.Delete(&model).Error
	return rst,err
}
