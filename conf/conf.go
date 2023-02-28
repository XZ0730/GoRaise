package conf

import (
	"Raising/dao"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPwd    string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string
	ConTent1   string
	ConTent2   string
	ConTent3   string

	Appmode string
	Port    string

	AccessKey   string
	SerectKey   string
	Bucket      string
	QiniuServer string

	AliAccessKey string
	AliSerectKey string
	SigName      string
	TemplateCode string

	Head      string
	AdminHead string
	EmailHead string

	Pagesize int

	Con = make(chan int, 1)
)

func Init() {
	file, err := ini.Load("./conf/conf.ini")
	fmt.Println("-------1111111111111111111111")
	if err != nil {
		fmt.Println("-------")
		panic(err)

	}
	fmt.Println("-------1111111111111111111111")
	fmt.Println("-------1111111111111111111111")
	Load(file)
	dbData := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	dao.Database(dbData)

}
func Load(file *ini.File) {
	LoadServiceConfig(file)
	LoadMysqlConfig(file)
	LoadRedisConfig(file)
	LoadValidEmailConfig(file)
	LoadQiniuConfig(file)
	LoadTokenConfig(file)
	LoadPageConfig(file)
	LoadAliyunConfig(file)
}
func LoadServiceConfig(file *ini.File) {
	Appmode = file.Section("service").Key("Appmode").String()
	Port = file.Section("service").Key("Port").String()
}
func LoadQiniuConfig(file *ini.File) {
	AccessKey = file.Section("Qiniu").Key("AccessKey").String()
	fmt.Println("AccessKey:", AccessKey)
	SerectKey = file.Section("Qiniu").Key("SerectKey").String()
	fmt.Println("SerectKey:", SerectKey)
	Bucket = file.Section("Qiniu").Key("Bucket").String()
	QiniuServer = file.Section("Qiniu").Key("QiniuServer").String()
}
func LoadMysqlConfig(file *ini.File) {
	DB = file.Section("mysql").Key("DB").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	fmt.Println("DbPassword =", DbPassword)
}
func LoadRedisConfig(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPwd = file.Section("redis").Key("RedisPwd").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()

}
func LoadValidEmailConfig(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
	ConTent1 = file.Section("email").Key("ConTent1").String()
	ConTent2 = file.Section("email").Key("ConTent1").String()
	ConTent3 = file.Section("email").Key("ConTent1").String()
}
func LoadTokenConfig(file *ini.File) {
	Head = file.Section("token").Key("head").String()
	AdminHead = file.Section("token").Key("adminHead").String()
	EmailHead = file.Section("token").Key("emailHead").String()
}
func LoadPageConfig(file *ini.File) {
	ps := file.Section("page").Key("pageSize").String()
	Pagesize, _ = strconv.Atoi(ps)
}
func LoadAliyunConfig(file *ini.File) {
	AliAccessKey = file.Section("Aliyun").Key("aliAccessKey").String()
	AliSerectKey = file.Section("Aliyun").Key("aliSerectKey").String()
	SigName = file.Section("Aliyun").Key("sigName").String()
	TemplateCode = file.Section("Aliyun").Key("tamplateCode").String()
}
