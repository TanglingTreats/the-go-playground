package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/TanglingTreats/go-cryptography-exp/contxt"
	crypt "github.com/TanglingTreats/go-cryptography-exp/crypto"
	"github.com/TanglingTreats/go-cryptography-exp/interfaces"
	iofile "github.com/TanglingTreats/go-cryptography-exp/io"
)

func main() {
	argsNum := len(os.Args)

	mode := flag.String("mode", "encrypt", "The mode to run the program in")
	flag.Parse()

	switch *mode {
	case "encrypt":
		if argsNum > 4 || argsNum <= 3 {
			fmt.Println("Only expecting one argument as a filename.")
			os.Exit(1)
		}
		args := os.Args[3:]
		filename := args[0]
		cryptoFun(filename)
	case "interface":
		interfaces.Exec()
	case "context":
		contxt.Exec()
	}

}

func cryptoFun(filename string) {
	// Open supplied filename
	file := iofile.ReadFile(filename)

	crypt.FreqAnalysis(string(file))
}
