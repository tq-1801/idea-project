package util

import (
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)
var SessionMap sync.Map
var CustomLogger *logrus.Logger
//var Claims cusfun.CustomClaims
func SysInit() {
	ReadConfig()
	CustomLogger = LoggerInit()
	DbConnect()
	RedisInit()
	KafkaProducerInit()
	InitEs()
	//FuncInit()

}

func DbConnect() {
	//ReadConfig()
	if Cfg.Db.Ip != "" {
		ok := Connect()
		if !ok {
			log.Println("db connect err!")
		} else {
			log.Println("db connected success!")
		}
	}

}
