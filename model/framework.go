package model

import (
	"strings"
	"path"
	"reflect"
)

type Framework struct {
	Name		string 	`form:"name" binding:"required"`
	Location 	string 	`form:"location" binding:"required"`
	Version 	string	`form:"version" binding:"required"`
	OS 			string	`form:"os" binding:"required"`
	Xcode		string 	`form:"xcode" binding:"required"`
}

func (f *Framework)ZipFilePath() string {
	return strings.Replace(f.Name, "-", "_", -1) + ".zip"
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