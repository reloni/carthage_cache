package utils

import (
	"buildben/carthage_cache/client/environment"
)

func LoadConfig() {
	environment.ServerAddress = "https://api.buildben.io/carthage"
}
