package utils

import (
	"os/exec"
	"bytes"
	"buildben/carthage_cache/client/web"
	"buildben/carthage_cache/client/model"
	"github.com/pierrre/archivefile/zip"
	"fmt"
	log "github.com/Sirupsen/logrus"
)



func HandleParsedFramework(f model.Framework) {
	switch actionForFramework(f) {
	case ActionDownload:
		fmt.Printf("%s found in Cloud!\n", f.Name)
		download(f)
		unarchive(f)
	case ActionCreate:
		fmt.Printf("%s doesn't exist in Cloud\n", f.Name)
		checkout(f)
		build(f)
		archive(f)
		upload(f)
	case ActionLocal:
		fmt.Printf("%s is not versioned - Cloud is unavaliable\n", f.Name)
		checkout(f)
		build(f)
	}
	fmt.Printf("%s up and ready.\n\n", f.Name)
}

type Action int

const (
	ActionDownload Action = iota
	ActionCreate
	ActionLocal
)

const githubLocation string = "github"

func actionForFramework(f model.Framework) Action{

	if exists(f) {
		return ActionDownload
	}

	if f.Location == githubLocation && len(f.Version) < 10 {
		return ActionCreate
	}

	return ActionLocal
}

func checkout(f model.Framework) {
	fmt.Printf("Checking out %s\n", f.Name)
	cmd := exec.Command("carthage", "checkout", f.Name, "--no-use-binaries")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func build(f model.Framework) {
	fmt.Printf("Building %s, it might take a while...\n", f.Name)
	cmd := exec.Command("carthage", "build", f.Name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func archive(f model.Framework) {
	fmt.Printf("Archiving %s\n", f.Name)
	zip.ArchiveFile(f.FilePath(), f.ZipFilePath(), nil)
}

func download(f model.Framework) {
	fmt.Printf("Downloading %s from Cloud\n", f.Name)
	web.Download(f)
}

func unarchive(f model.Framework) {
	fmt.Printf("Unzipping %s\n", f.Name)
	err := zip.UnarchiveFile(f.ZipFilePath(), f.Directory(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func upload(f model.Framework) {
	fmt.Printf("Uploading %s to Cloud\n", f.Name)
	web.Upload(f)
}

func exists(f model.Framework) bool {
	return web.Exists(f)
}