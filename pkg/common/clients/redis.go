package clients

import (
	"dumpapp_server/pkg/config"
	"github.com/go-redis/redis/v8"
)

var DumpRedis redis.Client

func init() {
	r := redis.NewClient(&redis.Options{
		Addr:     config.DumpConfig.AppConfig.Redis.Addr,
		Password: config.DumpConfig.AppConfig.Redis.Password, // no password set
		DB:       0,                                          // use default DB
	})
	DumpRedis = *r
}
