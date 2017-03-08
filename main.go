package main

import (
	"fmt"
	"os"

	"./fileutils"
	"./mysizes"
)

func main() {

	byteslice := make([]byte, 0) // size and capacity is 0

	// The append built-in function appends elements to the end of a slice.
	// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
	// If it does not, a new underlying array will be allocated.
	// Append returns the updated slice.
	for i := 0; i < 1; i++ {
		byteslice = append(byteslice, fileutils.FileToByteslice("files/pg100.txt")...)
	}
	fmt.Println(" ----- Playing with files / file descriptors ------ ")
	size := len(byteslice)
	//fmt.Println(size)
	fmt.Println(" ----- Size of the file ------- ")
	bs := mysizes.ByteSize(size)
	fmt.Println(bs)
	bs = mysizes.ByteSize(1048576)
	fmt.Println(bs)
	empty := fileutils.CreateEmptyFile("empty.txt")
	fmt.Printf(" ----- Created and empty file %s ----- \n", empty.Name())

	text1slice := fileutils.FileToByteslice("files/text1.txt")
	text2slice := fileutils.FileToByteslice("files/text2.txt")

	fmt.Println(text1slice)
	fmt.Println(text2slice)

	fileinfo := fileutils.GetFileInfo("files/text1.txt")
	mode := fileinfo.Mode()
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Size())
	fmt.Println(mode)
	fmt.Println(fileinfo.ModTime())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.Sys())
	fmt.Println(os.ModeSymlink)
	fmt.Println(os.ModeSetuid)
	fmt.Println(mode & os.ModeSymlink)
	fmt.Println(mode & os.ModeSetuid)

}
