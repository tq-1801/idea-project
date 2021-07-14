package util

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"time"
)

func LoggerInit() *logrus.Logger {
	logFilePath := ""
	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" { //开发模式
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
	} else { //部署模式
		logFilePath = "/opt/dsp/dsp/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}

	logFileName := Cfg.Logs.LogFile

	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	src.WriteString("\r\n")
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	//logger.SetFormatter(&logrus.JSONFormatter{})
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(fileName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		//logrus.FatalLevel: logWriter,
		//logrus.DebugLevel: logWriter,
		//logrus.WarnLevel:  logWriter,
		//logrus.ErrorLevel: logWriter,
		//logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		//TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return logger
}
