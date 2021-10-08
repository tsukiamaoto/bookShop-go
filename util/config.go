package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DATABASE_DRIVER"`
	DBSource      string `mapstructure:"DATABASE_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic("讀取設定檔出現錯誤，錯誤的原因為" + err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
