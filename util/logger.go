package util

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> fd910d7 (golang)
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

<<<<<<< HEAD
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
=======
type LastDay struct {
	Month int
	day   int
}

var (
	LogrusObj   *logrus.Logger
	lastDay     *LastDay
	Path        string
	logFileName string
)

func Init() {
	lastDay = new(LastDay)
	err := setDir()
	if err != nil {
		panic(err)
	}
	src := setoutToFile(Path)

	//实例化
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel) //设置日志级别
	logger.Out = src                   //设置输出
>>>>>>> fd910d7 (golang)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}
<<<<<<< HEAD
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
=======
func ReLogrusObj(Path string) *logrus.Logger {
	src := setoutToFile(Path)
	if LogrusObj != nil {
		LogrusObj.Out = src //设置输出
	}
	return LogrusObj
}

// setoutToFile方法返回logursobj 方法中比较日期，如果日期不同则直接创建新文件
func setoutToFile(Path string) *os.File {
	now := time.Now()

	if int(now.Month()) != lastDay.Month || now.Day() != lastDay.day {
		logFileName = now.Format("2006-01-02") + ".log"
	}
	fileName := path.Join(Path, logFileName) //日期不同新建文件

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Println("[log err]", err.Error())
	}
	return src
}
func setDir() error {
	now := time.Now()
	logFilePath := ""
	dir := "D:/Golang/Raising"
	logFilePath = dir + "/logs/"

	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {

			log.Println(err.Error())
			return err
		}
	}

	lastDay.Month = int(now.Month())
	lastDay.day = now.Day()
	//日志文件

	_, err = os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, os.ModePerm); err != nil {
			log.Println("[log err]", err.Error())
			return err
		}
	}
	//写入日志
	Path = logFilePath
	logFileName = now.Format("2006-01-02") + ".log"
	return nil
>>>>>>> fd910d7 (golang)
}
