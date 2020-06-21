package test

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func createMultipartFormData(personID []byte, image *os.File) (bytes.Buffer, *multipart.Writer) {
	var (
		b   bytes.Buffer
		err error
		w   = multipart.NewWriter(&b)
		fw  io.Writer
	)

	if fw, err = w.CreateFormField("personID"); err != nil {
		log.Fatal("Error creating writer", err)
	}
	if _, err = io.Copy(fw, bytes.NewReader(personID)); err != nil {
		log.Fatal("Error with io.Copy", err)
	}

	if fw, err = w.CreateFormFile("file", "test.jpg"); err != nil {
		log.Fatal("Error creating writer", err)
	}
	if _, err = io.Copy(fw, image); err != nil {
		log.Fatal("Error with io.Copy", err)
	}
	w.Close()
	image.Seek(0, 0)
	return b, w
}

func recognizeMultipartFormData(keys []string, image *os.File) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var w = multipart.NewWriter(&b)

	fw, err := w.CreateFormField("photoIDs")
	if err != nil {
		log.Fatal("Error creating writer", err)
	}

	if _, err = io.Copy(fw, bytes.NewReader([]byte(strings.Join(keys, ",")))); err != nil {
		log.Fatal("Error with io.Copy", err)
	}
	if fw, err = w.CreateFormFile("file", "test.jpg"); err != nil {
		log.Fatal("Error creating writer", err)
	}

	if _, err = io.Copy(fw, image); err != nil {
		log.Fatal("Error with io.Copy", err)
	}
	w.Close()
	image.Seek(0, 0)
	return b, w
}

func LogTestCase(method string, path string, useCase string, expect string) {
	log.WithFields(log.Fields{
		"METHOD": method,
		"PATH": path,
		"USE_CASE": useCase,
		"EXPECT": expect,
	}).Info()
}