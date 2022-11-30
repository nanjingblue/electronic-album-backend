package cache

import (
	"electronic-album/global"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func SetupRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     global.DatabaseSetting.RedisAddr,
		Password: global.DatabaseSetting.RedisPassword,
		DB:       global.DatabaseSetting.RedisDB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	RedisClient = client
	return err
}
