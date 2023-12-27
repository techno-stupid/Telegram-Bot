package app

import (
	"github.com/spf13/viper"
)

func Init() {
	loadEnvVariables()
}
func loadEnvVariables() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
