package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := serOutPutFile()
	if LogrusObj != nil {

		LogrusObj.Out = src
		return
	}
	//实例化
	fmt.Println("1111111111111111111111")
	logger := logrus.New()
	logger.Out = src                   //设置输出
	logger.SetLevel(logrus.DebugLevel) //设置日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}
func serOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err != nil {
		logFilePath = dir + "/logs/"
	}
	fmt.Println(logFilePath)
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			fmt.Println("12345")
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fileName, 0777); err != nil {
			fmt.Println("123456")
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入日志
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("123456")
		return nil, err
	}
	return src, nil
}
