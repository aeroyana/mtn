package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// uploads a file via form submit to the given URL
func uploadFile(filename string) error {
	if viper.GetBool("upload") && viper.GetString("upload_url") != "http://example.com/upload" {
		log.Infof("uploading file...")
		targetUrl := viper.GetString("upload_url")
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("image", filename)
		if err != nil {
			log.Errorf("error writing to buffer")
			return err
		}

		// open file handle
		fh, err := os.Open(filename)
		if err != nil {
			log.Errorf("error opening file: %s", err)
			return err
		}

		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			log.Errorf("error iocopy: %s", err)
			return err
		}

		contentType := bodyWriter.FormDataContentType()
		bodyWriter.Close()

		resp, err := http.Post(targetUrl, contentType, bodyBuf)
		if err != nil {
			log.Errorf("error sending request to server: %s", err)
			return err
		}
		defer resp.Body.Close()
		resp_body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("error reading server response: %s", err)
			return err
		}
		log.Infof("Server response:\n%s", string(resp_body))
		return nil
	}
	return nil
}
