package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPwd    string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/conf.ini")
	if err != nil {
		panic(err)
	}
	LoadRedisConfig(file)
	Redis()
}
func LoadRedisConfig(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPwd = file.Section("redis").Key("RedisPwd").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()

}
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPwd,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(RedisPwd, ":", RedisAddr)
		panic(err)
	}
	RedisClient = client
}
