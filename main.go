package main

import (
	"fmt"
	"os"

	crypt "github.com/TanglingTreats/go-cryptography-exp/crypto"
	iofile "github.com/TanglingTreats/go-cryptography-exp/io"
)

func main() {
	argsNum := len(os.Args)

	if argsNum >= 3 || argsNum < 2 {
		fmt.Println("Only expecting one argument as a filename.")
		os.Exit(1)
	}

	args := os.Args[1:]
	filename := args[0]

	// Open supplied filename
	file := iofile.ReadFile(filename)

	crypt.FreqAnalysis(string(file))
}
