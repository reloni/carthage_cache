package main

import (
	. "buildben/carthage_cache/client/utils"
)

func main() {
	LoadConfig()
	LoadProcessArguments()

	ParseCartfile(func (line string) {
		f := FrameworkFromOgdlString(line)
		HandleParsedFramework(f)
	})
}

