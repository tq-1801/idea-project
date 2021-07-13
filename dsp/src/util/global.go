package util

import (
	"log"
)

//var Claims cusfun.CustomClaims
func SysInit() {

	ReadConfig()
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
