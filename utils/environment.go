package utils

import (
	"flag"
	"github.com/joho/godotenv"
	"buildben/carthage_cache/client/environment"
)

func LoadProcessArguments() {
	flag.StringVar(&environment.Platform,"platform", "iOS", "Platform. Default is iOS.")
	flag.Parse()
}

func LoadEnvDefaults() {
	godotenv.Load("fastlane/.env.default")
}

