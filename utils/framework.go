package utils

import (
	"os/exec"
	"bytes"
	"buildben/carthage_cache/client/web"
	"buildben/carthage_cache/client/model"
	"github.com/pierrre/archivefile/zip"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"regexp"
	"bufio"
	"os"
	"path/filepath"
	"path"
	"github.com/jhoonb/archivex"
	"github.com/rveen/ogdl"
	"strings"
	"buildben/carthage_cache/client/environment"
)


func FrameworkFromOgdlString(s string) model.Framework {
	var f model.Framework = model.Framework{}

	g := ogdl.ParseString(s)
	host :=  g.GetAt(0)
	name := host.GetAt(0)
	fVersion := name.GetAt(0)


	f.Location = host.String()
	f.Name = strings.Split(name.String(), "/")[1]
	f.Version = fVersion.String()
	f.OS = environment.Platform
	f.Xcode = XcodeBuildVersion()

	return f
}

func HandleParsedFramework(f model.Framework) {
	switch actionForFramework(f) {
	case ActionDownload:
		fmt.Printf("%s found in Cloud!\n", f.Name)
		download(f)
		unarchive(f)
	case ActionCreate:
		fmt.Printf("%s doesn't exist in Cloud.\n", f.Name)
		checkout(f)
		n := build(f)
		if len(n) == 0 { return }
		archive(f, n)
		upload(f)
	case ActionLocal:
		fmt.Printf("%s is not versioned - Cloud is unavaliable.\n", f.Name)
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

	if f.Location == githubLocation && len(f.Version) < 20 {
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

func build(f model.Framework) []string {
	fmt.Printf("Building %s, it might take a while...\n", f.Name)
	cmd := exec.Command("carthage", "build", f.Name, "--platform", f.OS)
	var out bytes.Buffer
	cmd.Stdout = &out
	
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	filePath := findLogFilePath(out.String())
	return findFrameworkFiles(filePath)
}


func findLogFilePath(inc string) string {
	r, _ := regexp.Compile("/var.*\\.log")
	return r.FindString(inc)
}

func findFrameworkFiles(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("/.*\\.framework$")

	files := []string {}
	for scanner.Scan() {
		line := scanner.Text()
		if result := r.FindString(line); result != "" {
			_, name := filepath.Split(result)
			files = appendIfMissing(files, name)
		}
	}
	if len(files) == 0 {
		fmt.Print("No schemes found for selected platform. Skipping...\n\n")
	}

	return files
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func archive(f model.Framework, paths []string) {
	fmt.Printf("Archiving %s\n", f.Name)
	archive := new(archivex.ZipFile)
	archive.Create(f.ZipFilePath())
	for _, p := range paths {
		archive.AddAll(path.Join(f.Directory(), p), true)
	}
	archive.Close()
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