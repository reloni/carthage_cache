package utils

import (
	"os/exec"
	. "strings"
	log "github.com/Sirupsen/logrus"
)


func XcodeBuildVersion() string {
	out, err := exec.Command("xcodebuild", "-version").CombinedOutput()
	if err != nil {
		log.Error(string(out))
		log.Fatal(err)
	}
	return Split(Split(string(out), "\n")[1], " ")[2]
}