package web


import (
	"fmt"
	"net/http"
	"buildben/carthage_cache/client/model"
	"os"
	"io"
	log "github.com/Sirupsen/logrus"
	"reflect"
	"strings"
	"buildben/carthage_cache/client/environment"
	"crypto/tls"
)


func Upload(f model.Framework) {
	req, err := fileUploadRequest(joinUrl(environment.ServerAddress, "upload"), f.Map(),"upload", f.ZipFilePath())
	if err != nil {
		log.Error(fmt.Sprintf("Unable to upload %s to the cloud ==> %s", f.Name, err.Error()))
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err = client.Do(req)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to upload %s to the cloud ==> %s", f.Name, err.Error()))
		return
	}
}

func Exists(f model.Framework) bool {
	req, _ := http.NewRequest("GET", joinUrl(environment.ServerAddress, "exists"), nil)
	appendGetFrameworkQuery(req, f)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	return res.StatusCode == http.StatusOK

}

func Download(f model.Framework) bool{
	req, _ := http.NewRequest("GET", joinUrl(environment.ServerAddress, "download"), nil)
	appendGetFrameworkQuery(req, f)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	out, err := os.Create(f.ZipFilePath())
	defer out.Close()


	defer res.Body.Close()
	_, err = io.Copy(out, res.Body)

	return !(res.StatusCode == http.StatusBadRequest || err != nil)
}


func appendGetFrameworkQuery(req *http.Request, f model.Framework) {
	q := req.URL.Query()

	val := reflect.ValueOf(&f).Elem()

	for i := 0; i < val.NumField(); i++ {
		q.Add(strings.ToLower(val.Type().Field(i).Name), val.Field(i).Interface().(string))
	}

	req.URL.RawQuery = q.Encode()
}