package fileutils

import (
	"log"
	"os"
)

func GetFileInfo(filename string) os.FileInfo {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fi
}
