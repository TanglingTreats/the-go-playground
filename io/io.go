package iofile

import (
	"fmt"
	"os"
)

// Return file as bytes with given file path as a string
func ReadFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return file
}
