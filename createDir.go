package logs

import (
	"log"
	"os"
)

func createDir() {
	// Verify if dir exists
	_, err := os.Stat(logFolder)

	// If negative, create
	if os.IsNotExist(err) {
		err := os.MkdirAll(logFolder, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
