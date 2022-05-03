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
	Databases     map[string]*Database
	ServerAddress string
	AllowOrigins  []string
	Redis         *Redis
}

type Database struct {
	Name   string
	Source string
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic("讀取設定檔出現錯誤，錯誤的原因為" + err.Error())
	}

	var dbs = make(map[string]*Database)
	dbs["shopCart"] = getDatabase("shopCart")
	dbs["default"] = getDatabase("default")

	serverAddress := fmt.Sprintf("%s:%d", viper.GetString("application.host"), viper.GetInt("application.port"))
	redis := &Redis{
		Address:  viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	allowOrigins := viper.GetStringSlice("application.cors.allowOrigins")

	config := &Config{
		Databases:     dbs,
		ServerAddress: serverAddress,
		Redis:         redis,
		AllowOrigins:  allowOrigins,
	}

	return config
}

func getDatabase(name string) *Database {
	dbName := viper.GetString(fmt.Sprintf("databases.%s.dbname", name))
	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Taipei",
		viper.GetString(fmt.Sprintf("databases.%s.host", name)),
		viper.GetInt(fmt.Sprintf("databases.%s.port", name)),
		viper.GetString(fmt.Sprintf("databases.%s.user", name)),
		viper.GetString(fmt.Sprintf("databases.%s.password", name)),
		viper.GetString(fmt.Sprintf("databases.%s.dbname", name)),
	)

	return &Database{
		Name:   dbName,
		Source: source,
	}
}
