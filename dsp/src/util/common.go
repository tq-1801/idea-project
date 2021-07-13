package util

import (
	"fmt"
	"net"
	"os/exec"
	"reflect"
	"runtime"

	"strconv"
	"strings"
)

/*
字符串数组去重
*/
func RemoveDuplicate(s []string) []string {
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func HashStr64(sep string) int64 {
	primeRK := uint64(1099511628211)
	hash := uint64(14695981039346656037)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint64(sep[i])
	}
	return int64(hash)
}

func HashStr32(sep string) int32 {
	primeRK := uint32(16777619)
	hash := uint32(2166136261)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	return int32(hash)
}

//点分十进制的ip转为int值
func InetStr2Number(ip string) uint32 {
	if !IsvalidIpv4(ip) {
		return 4294967295 //uint 32 的最大值，255.255.255.255
	}
	str := strings.Split(ip, ".")
	high2432, _ := strconv.Atoi(str[0])
	high1624, _ := strconv.Atoi(str[1])
	high816, _ := strconv.Atoi(str[2])
	high08, _ := strconv.Atoi(str[3])
	result := (high2432 << 24) + (high1624 << 16) + (high816 << 8) + high08
	return uint32(result)
}

//判断字符串是否为有效的ip
func IsvalidIpv4(ip string) bool {
	ipaddr := net.ParseIP(ip)
	if ipaddr == nil {
		return false
	}
	if ipaddr.To4() == nil {
		return false
	}
	return true
}

//将IP地址转化为二进制String
func Ip2binary(ip string) string {
	str := strings.Split(ip, ".")
	var ipstr string
	for _, s := range str {
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		ipstr = ipstr + fmt.Sprintf("%08b", i)
	}
	return ipstr
}

//判断字符串是否数字（包括浮点数）
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

//判断字符串是否整数
func IsNum(s string) bool {
	_, err := strconv.ParseInt(s, 10, 32)
	return err == nil
}

/*
interface（string）类型转换为字符串数组类型
*/
func InterfaceListToStrList(list []interface{}) (res []string) {
	res = make([]string, len(list))
	for i, v := range list {
		res[i] = v.(string)
	}
	return res
}

/*
interface类型转换为int数组类型
*/
func InterfaceListToIntList(list []interface{}) (res []int) {
	res = make([]int, len(list))
	for i, v := range list {
		var n int
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			n, _ = strconv.Atoi(v.(string))
		case reflect.Float64:
			n = int(v.(float64))
		case reflect.Bool:
			if v.(bool) {
				n = 1
			} else {
				n = 0
			}
		}
		res[i] = n
	}
	return res
}

/*
interface（float）类型转换为int数组类型
*/
func InterfaceIntListToIntList(list []interface{}) (res []int) {
	res = make([]int, len(list))
	for i, v := range list {
		//n,_ := strconv.Atoi(v.(int))
		res[i] = int(v.(float64))
	}
	return res
}

//执行linux命令
func RunInLinux(cmd string) error {
	fmt.Println("Running Linux cmd:", cmd)
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

//执行linux命令,并等待执行完毕
func RunInLinux2(cmd string) error {
	fmt.Println("Running Linux cmd:", cmd)
	c := exec.Command("bash", "-c", cmd)
	err := c.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

//执行linux命令并获取结果
func RunInLinux3(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:", cmd)
	b, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(b), err
}

//执行windows命令
func RunInWindows(cmd string) error {
	fmt.Println("Running Win cmd:", cmd)
	_, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

//执行windows命令并获取结果
func RunInWindows3(cmd string) (string, error) {
	fmt.Println("Running Win cmd:", cmd)
	b, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println("Running Win res:", string(b))
	return string(b), err
}

/*
服务重启
*/
func SvcRestart() (err error) {
	wincmd := "ipconfig"
	liuxstopcmd := "systemctl stop nfca"
	liuxstartcmd := "systemctl start nfca"
	linuxKillcmd1 := "ps -ef |grep nmap |awk  '{print $2}' |xargs kill -9"
	linuxKillcmd2 := "ps -ef |ps -ef |grep masscan |awk  '{print $2}' |xargs kill -9"
	if runtime.GOOS == "windows" {
		err = RunInWindows(wincmd)
	} else {
		err = RunInLinux2(liuxstopcmd)
		_ = RunInLinux2(linuxKillcmd1)
		_ = RunInLinux2(linuxKillcmd2)
		err = RunInLinux2(liuxstartcmd)
	}
	return err
}

/*
服务重启
*/
func SwitchAclQuery() (err error) {
	wincmd := "ipconfig"
	liuxstopcmd := "systemctl stop nfca"
	liuxstartcmd := "systemctl start nfca"
	linuxKillcmd1 := "ps -ef |grep nmap |awk  '{print $2}' |xargs kill -9"
	linuxKillcmd2 := "ps -ef |ps -ef |grep masscan |awk  '{print $2}' |xargs kill -9"
	if runtime.GOOS == "windows" {
		err = RunInWindows(wincmd)
	} else {
		err = RunInLinux(liuxstopcmd)
		_ = RunInLinux(linuxKillcmd1)
		_ = RunInLinux(linuxKillcmd2)
		err = RunInLinux(liuxstartcmd)
	}
	return nil
}

//func InsertLog(user, ip, operType string, logType string, content string, msg string, err error) {
//	id := int(HashStr32(uuid.NewV4().String()))
//	log := model.Syslog{
//		Id:             id,
//		Operator:       user,
//		Opertype:       operType,
//		Result:         msg,
//		Opertime:       time.Now(),
//		Content:        content,
//		Address:        ip,
//		Keywords:       "",
//		Logmillisecond: time.Now().UnixNano() / 1e6,
//		Logtype:        logType,
//	}
//	s := user + "," + log.Address + "," + content + "," + msg + "," + logType + "," + log.Opertime.Format("2006-01-02 15:04:05")
//	log.Keywords = s
//	//入本地库
//	DbConn.Create(&log)
//
//}
