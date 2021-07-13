package util

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

func Time2Nano(timeString string) int64 {
	times, _ := time.Parse(time.RFC3339, timeString)
	return times.UnixNano() / 1e6
}

/*
 * 判断是否是虚拟机
 */
//func IsVirtualMachine() (bool, string) {
//	system_info, e := ExecShell("dmidecode -s system-manufacturer|awk '{print $1}'")
//	if e != nil {
//		log.Println(e)
//		return false, ""
//	}
//	system_info = strings.Trim(system_info, "\n")
//	system_info = strings.Trim(system_info, ",")
//	if (system_info == "VMware") || (system_info == "Xen") {
//		return true, system_info
//	}
//	return false, system_info
//}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsFileDirExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func InitLogger(logfile string) error {
	log.SetOutput(&lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
	log.SetFlags(log.Ldate | log.Ltime | log.LstdFlags | log.Lshortfile)
	return nil
}
