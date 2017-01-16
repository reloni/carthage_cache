package main

import (
	"buildben/carthage_cache/client/model"
	"buildben/carthage_cache/client/utils"
)

func main() {
	utils.LoadConfig()

	utils.LoadEnvDefaults()
	utils.LoadProcessArguments()

	utils.ParseCartfile(func (line string) {
		f := model.FrameworkFromOgdlString(line)
		utils.HandleParsedFramework(f)
	})
}

