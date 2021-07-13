package util

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"runtime"
)

type Config struct {
	Db struct {
		Ip           string `json:"ip"`
		Port         string `json:"port"`
		Name         string `json:"name"`
		User         string `json:"user"`
		Password     string `json:"password"`
		RootPassword string `json:"rootPassword"`
	}
	Redis struct {
		RedisIps          string `json:"redisIps"`
		RedisPwd          string `json:"redisPwd"`
		SessionExpMinutes int64  `json:"sessionExpMinutes"`
	}
	Es struct {
		EsIps string `json:"esIps"`
	}
	Kafka struct {
		Urls string `json:"urls"`
	}

	FtpPath struct {
		SoftwarePath string `json:"softwarePath"`
	}
	Service struct {
		Names     string `json:"names"`
		AutoStart bool   `json:"autoStart"`
	}

	Scan struct {
		StatusScan uint64 `json:"statusScan"`
	}
	Logs struct {
		LogFile string `json:"logFile"`
	}
	Mode struct {
		WorkMode   string `json:"workMode"`
		Master     string `json:"master"`
		ManageIpv4 string `json:"manageIpv4"`
		ManageIpv6 string `json:"manageIpv6"`
		Vip        string `json:"vip"`
		SlaveIps   string `json:"slaveIps"`
	}
	Nets struct {
		ManageEth string `json:"manageEth"`
		FixedEth  string `json:"fixedEth"`
	}

	IptPath struct {
		IptRules string `json:"iptRules"`
	}

	SysInfo struct {
		Version    string `json:"version"`
		NewVersion string `json:"newVersion"`
	}
}

var Cfg Config

var fileName = "src/config/config.ini"

//读取配置文件到内存
func ReadConfig() bool {
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { //开发模式
		fileName = "src/config/config.ini"
	} else { //部署模式
		fileName = "/opt/somf/somf/conf/config.ini"
	}
	err := gcfg.ReadFileInto(&Cfg, fileName)
	if err != nil {
		log.Println(err)
		return false
	}

	fmt.Println("db:")
	fmt.Println(Cfg.Db)

	var aes *AesEncrypt
	Cfg.Db.Password, _ = aes.Decrypt(Cfg.Db.Password)
	Cfg.Db.RootPassword, _ = aes.Decrypt(Cfg.Db.RootPassword)
	Cfg.Redis.RedisPwd, _ = aes.Decrypt(Cfg.Redis.RedisPwd)

	fmt.Println("scan:")
	fmt.Println(Cfg.Scan)

	fmt.Println("es:")
	fmt.Println(Cfg.Es)

	fmt.Println("iptPath:")
	fmt.Println(Cfg.IptPath)

	return true
}

func SetConfig(section, key, value string) bool {
	cfg, err := ini.Load(fileName)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	//修改某个值然后进行保存
	cfg.Section(section).Key(key).SetValue(value)
	cfg.SaveTo(fileName)
	ReadConfig()
	return true
}
