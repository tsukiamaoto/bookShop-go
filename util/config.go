package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string 
	ServerAddress string 
}

func LoadConfig() *Config {
	viper.AddConfigPath("./config")
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

	config := &Config{
		DBSource: dsn,
		ServerAddress: serverAddress,
	}

	return config
}
