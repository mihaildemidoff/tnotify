package service

import (
	"io/ioutil"
	"log"
	"os"
)

func WriteTextToFile(text string) string {
	file, err := ioutil.TempFile(os.TempDir(), "tnotify-*.txt")
	if err != nil {
		log.Fatalf("Error occurred during temp file creation: %s", err)
	}
	defer file.Close()
	if _, err = file.Write([]byte(text)); err != nil {
		log.Fatalf("Error occurred during writing to temp file: %s", err)
	}

	fileName := file.Name()
	return fileName
}
