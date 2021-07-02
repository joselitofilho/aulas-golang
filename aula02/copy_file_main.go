package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() // Correct!

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close() // Correct!

	written, err = io.Copy(dst, src)
	// src.Close() // Wrong!
	// dst.Close() // Wrong!
	return
}

func main() {
	written, err := CopyFile("myFile.txt", "copyFile.txt")
	fmt.Println(written, err)
}
