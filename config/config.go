package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Redis struct {
	Address  string
	Password string
	DB       int
}

type Config struct {
	DBSource      string
	ServerAddress string
	SessionKey    string
	Redis         *Redis
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic("讀取設定檔出現錯誤，錯誤的原因為" + err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%d dbname=%s sslmode=disable",
		viper.GetString("database.host"), viper.GetInt("database.port"), viper.GetString("database.user"),
		viper.GetInt("database.password"), viper.GetString("database.dbname"))
	serverAddress := fmt.Sprintf("%s:%d", viper.GetString("test.host"), viper.GetInt("test.port"))
	redis := &Redis{
		Address:  viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	sessionKey := viper.GetString("test.sessionKey")

	config := &Config{
		DBSource:      dsn,
		ServerAddress: serverAddress,
		Redis:         redis,
		SessionKey:    sessionKey,
	}

	return config
}
