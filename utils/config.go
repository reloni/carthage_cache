package utils

import (
	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
	"buildben/carthage_cache/client/environment"
)



func LoadConfig() {
	viper.SetConfigName("carthage")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.buildben/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	reloadVariables()
}

func reloadVariables() {
	environment.ServerAddress = viper.GetString("server_address")
}