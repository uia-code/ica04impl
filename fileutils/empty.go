package fileutils

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func CreateEmptyFile(filename string) os.FileInfo {
	newFile, err = os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	fileinfo, _ := newFile.Stat()
	defer newFile.Close()
	return fileinfo
}
