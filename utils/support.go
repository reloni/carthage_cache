package utils

import (
	"io/ioutil"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var carthageBuildStaticScriptName = "carthage_build_static.sh"
var pythonLdScriptName = "ld.py"

func ExportBuildStaticScripts() (string, string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	dirName := fmt.Sprintf("/tmp/%s_%d/", "bb_carthage_scripts", r1.Intn(10000))
	ldData, _ := Asset("resources/" + pythonLdScriptName)
	shData, _ := Asset("resources/" + carthageBuildStaticScriptName)

	os.MkdirAll(dirName, os.ModePerm)
	ioutil.WriteFile(dirName + pythonLdScriptName, ldData , 0744)
	ioutil.WriteFile(dirName + carthageBuildStaticScriptName, shData , 0744)

	return dirName + pythonLdScriptName, dirName + carthageBuildStaticScriptName
}