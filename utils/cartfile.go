package utils

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"bufio"
)

type callback func(string)

func ParseCartfile(cb callback) {
	file, err := os.Open("Cartfile.resolved")
	if err != nil {
		log.Fatal("Cartfile.resolved doesn't exists, aborting.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cb(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}