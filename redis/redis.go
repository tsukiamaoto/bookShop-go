package redis

import (
	"context"
	"time"

	Config "github.com/tsukiamaoto/bookShop-go/config"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var RDB *redis.Client
var ctx = context.Background()

func newRDB() *redis.Client {
	config := Config.LoadConfig()

	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Address,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
}

func ConnectRDB() {
	RDB = newRDB()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Error(err)
	}
}

func GetRDB() *redis.Client {
	if RDB != nil {
		return RDB
	}
	return newRDB()
}

func Get(key string) string {
	value, err := GetRDB().Get(ctx, key).Result()
	if err != nil {
		log.Error(err)
	}
	return value
}

func Set(key string, value interface{}) {
	_, err := GetRDB().Set(ctx, key, value, 0).Result()
	if err != nil {
		log.Error(err)
	}
}

func GetEx(key string, expiration time.Duration) string {
	value, err := GetRDB().GetEx(ctx, key, expiration).Result()
	if err != nil {
		log.Error(err)
	}
	return value
}

func SetEx(key string, value interface{}, expiration time.Duration) {
	_, err := GetRDB().Set(ctx, key, value, expiration).Result()
	if err != nil {
		log.Error(err)
	}
}

func Exists(key string) bool {
	ok, err := GetRDB().Exists(ctx, key).Result()
	if err != nil {
		log.Error(err)
	}

	return ok == 1
}
