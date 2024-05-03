package main

import (
	"flag"
	"fmt"
	"os"

	crypt "github.com/TanglingTreats/go-cryptography-exp/crypto"
	iofile "github.com/TanglingTreats/go-cryptography-exp/io"
)

type Player interface {
	KickBall() int
}

type FootballPlayer struct {
	speed   int
	stamina int
}

func (fp FootballPlayer) KickBall() int {
	return fp.speed * fp.stamina
}

func main() {
	argsNum := len(os.Args)

	mode := flag.String("mode", "encrypt", "The mode to run the program in")

	switch *mode {
	case "encrypt":
		fmt.Println(argsNum)
		if argsNum > 4 || argsNum <= 3 {
			fmt.Println("Only expecting one argument as a filename.")
			os.Exit(1)
		}
		args := os.Args[3:]
		filename := args[0]
		cryptoFun(filename)

	}

}

func cryptoFun(filename string) {
	// Open supplied filename
	file := iofile.ReadFile(filename)

	crypt.FreqAnalysis(string(file))
}
