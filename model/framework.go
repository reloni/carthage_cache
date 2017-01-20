package model

import (
	"strings"
	"path"
	"github.com/rveen/ogdl"
	"reflect"
	"buildben/carthage_cache/client/environment"
	"os"
)

type Framework struct {
	Name		string 	`form:"name" binding:"required"`
	Location 	string 	`form:"location" binding:"required"`
	Version 	string	`form:"version" binding:"required"`
	OS 		string	`form:"os" binding:"required"`
	Xcode		string 	`form:"xcode" binding:"required"`
}


const Xcode_Version_Key string = "XCODE_VERSION"

func (f *Framework)ZipFilePath() string {
	return strings.Replace(f.Name, "-", "_", -1) + ".framework.zip"
}
func (f *Framework)FilePath() string {
	return path.Join(f.Directory(), strings.Replace(f.Name, "-", "_", -1) + ".framework")
}

func (f *Framework)Directory() string {
	return path.Join("Carthage", "Build", f.OS)
}

func (f *Framework)Map() map[string]string {
	var dictionary = make(map[string]string)

	val := reflect.ValueOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		dictionary[strings.ToLower(val.Type().Field(i).Name)] = val.Field(i).Interface().(string)
	}

	return dictionary
}


func FrameworkFromOgdlString(s string) Framework {
	var f Framework = Framework{}

	g := ogdl.ParseString(s)
	host :=  g.GetAt(0)
	name := host.GetAt(0)
	fVersion := name.GetAt(0)


	f.Location = host.String()
	f.Name = strings.Split(name.String(), "/")[1]
	f.Version = fVersion.String()
	f.OS = environment.Platform
	f.Xcode = os.Getenv(Xcode_Version_Key)

	return f
}
