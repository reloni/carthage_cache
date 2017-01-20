package utils

import (
	//"github.com/spf13/viper"
	//log "github.com/Sirupsen/logrus"
	"buildben/carthage_cache/client/environment"
)



func LoadConfig() {
	environment.ServerAddress = "https://api.buildben.io/carthage"
	//viper.SetConfigName("carthage")
	//viper.SetConfigType("json")
	//viper.AddConfigPath("$HOME/.buildben/")
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatal(err)
	//}
	//reloadVariables()
}

//func reloadVariables() {
//	environment.ServerAddress = viper.Get("api_server_address")
//	log.Info(environment.ServerAddress)
//}