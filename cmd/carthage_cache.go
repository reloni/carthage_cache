package main

import (
	. "buildben/carthage_cache/client/utils"
	"buildben/carthage_cache/client/environment"
)

func main() {
	LoadConfig()
	LoadProcessArguments()

	ParseCartfile(func (line string) {
		f := FrameworkFromOgdlString(line)
		if environment.BuildStatic {  f.Linking = "static" }
		HandleParsedFramework(f)
	})
}

