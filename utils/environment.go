package utils

import (
	"flag"
	"buildben/carthage_cache/client/environment"
)

func LoadProcessArguments() {
	flag.StringVar(&environment.Platform,"platform", "iOS", "Platform. Default is iOS.")
	flag.BoolVar(&environment.BuildStatic,"static", false, "Should be static frameworks. Default is NO.")
	flag.Parse()
}
