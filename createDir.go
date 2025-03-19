package logs

import (
	"log"
	"os"
)

func createDir() {
	// Verify if dir exists
	_, err := os.Stat(folderLogs)

	// If negative, create
	if os.IsNotExist(err) {
		err := os.MkdirAll(folderLogs, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
