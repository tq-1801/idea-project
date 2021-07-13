/**
 * @Author: mxp
 * @Description:
 * @File:  connectdb
 * @Version: 1.0.0
 * @Date: 2021/3/16 11:13
 */

package util

import (
	"fmt"
	"github.com/ttjio/gorm-logrus"

	//"github.com/ttjio/gorm-logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

var ApproveMap sync.Map
var DbConn = new(gorm.DB) //全局的数据库连接使用

/*
*初始化数据库连接，系统完成后，必须关闭defer DbConn.Close()
 */
func Connect() (b bool) {
	var err error
	dbLogger := LoggerInit()
	for i := 1; i <= 10; i++ {
		dsn := Cfg.Db.User + ":" + Cfg.Db.Password + "@tcp(" + Cfg.Db.Ip + ":" + Cfg.Db.Port + ")/" + Cfg.Db.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		DbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gorm_logrus.New(dbLogger),
			//Logger: newLogger,
		})
		if err != nil {
			log.Println("连接数据库失败,尝试重连……")
		} else {
			sqlDB, err1 := DbConn.DB()
			if err1 != nil {
				fmt.Println(err1.Error())
			}
			// SetMaxIdleConns 设置空闲连接池中连接的最大数量
			sqlDB.SetMaxIdleConns(10)

			// SetMaxOpenConns 设置打开数据库连接的最大数量
			sqlDB.SetMaxOpenConns(100)

			// SetConnMaxLifetime 设置了连接可复用的最大时间
			sqlDB.SetConnMaxLifetime(time.Hour)
			break
		}

		if i == 10 {
			log.Println("连接数据库失败！")
			return false
		}
		time.Sleep(6 * time.Second)
	}
	return true
}


