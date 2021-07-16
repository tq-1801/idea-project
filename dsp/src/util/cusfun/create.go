package cusfun

import (
	"gorm.io/gorm"
)

// 增加
func GetSqlByCreate(db *gorm.DB, model interface{}) (rst interface{},err error) {
	err = db.Create(&model).Error
	return rst,err

}
