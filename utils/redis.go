package utils

import (
	"encoding/json"
	"time"

	"github.com/beego/beego"
	"github.com/beego/beego/cache"
	_ "github.com/beego/beego/cache/redis"
	"github.com/beego/beego/v2/core/logs"
)

var (
	Redis        cache.Cache
	RedisTimeout time.Duration //默认过期时间1天
)

func init() {
	//redis初始化配置
	var err error
	config := getRedisConfig()
	Redis, err = cache.NewCache("redis", config)
	if err != nil {
		logs.Error("连接redis出错", err)
	}
	RedisTimeout = 86400 * time.Second
}

func getRedisConfig() string {
	key := beego.AppConfig.String("redis_key") //redis key前缀 例 synolux:xxx
	redis_host := beego.AppConfig.String("redis_host")
	redis_password := beego.AppConfig.String("redis_password")
	redis_port := beego.AppConfig.String("redis_port")
	redis_db := beego.AppConfig.String("redis_db")
	redisHash := make(map[string]interface{})
	redisHash["key"] = key
	redisHash["conn"] = redis_host + ":" + redis_port
	redisHash["dbNum"] = redis_db
	redisHash["password"] = redis_password
	redisConfig, _ := json.Marshal(redisHash)
	return string(redisConfig)
}
