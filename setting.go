package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func setConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetDefault("application.port", 9999)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic("讀取設定檔出現錯誤，錯誤的原因為" + err.Error())
	}
	fmt.Printf("application port:" + viper.GetString("application.port"))
}
