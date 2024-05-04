package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/TanglingTreats/the-go-playground/contxt"
	crypt "github.com/TanglingTreats/the-go-playground/crypto"
	"github.com/TanglingTreats/the-go-playground/interfaces"
	iofile "github.com/TanglingTreats/the-go-playground/io"
)

/**
 * Idiomatic Go
 */
// ---- Declaring constants ----
const (
	Name    = "The Go Playground"
	Version = "1.0.0"
)

// ---- Variable grouping ----
var (
	x = 1
	y = 1
)

// ---- enums ----
type Suit byte

// prefix with enum name
const (
	SuitSpades Suit = iota
	SuitHearts
	SuitDiamonds
	SuitClubs
)

// ---- struct initialization ----
type Position struct {
	X int
	Y int
}

// ---- Constructor
// Use 'New' prefix for constructor functions
// If in a package, omit type name. E.g. New()
func NewPosition(x, y int) Position {
	// Use named variables to initialize struct
	return Position{
		X: x,
		Y: y,
	}
}

// ---- interface declaration and naming ----
// Use 'er' suffix for interface
type Reader interface {
	Read() string
}

// compose interfaces
type Writer interface {
	Get() string
}

type ReadWriter interface {
	Reader
	Writer
}

// ---- mutex grouping ----
type Rotation struct {
	X int
	Y int
	Z int

	WMutex sync.RWMutex
	W      int
}

// Functions that panic; Prefix with 'Must'
func MustParseIntFromString(s string) int {
	// Logic
	if len(s) == 0 {
		panic("Not implemented")
	}

	return int(s[0] - '0')
}

// ---- Function grouping ----
// Arrange exported functions first. Simpler functions come after
func (r *Rotation) Rotate(x, y, z int) {
	r.WMutex.Lock()
	defer r.WMutex.Unlock()

	r.X = x
	r.Y = y
	r.Z = z
}

func (r *Rotation) Get() int {
	r.WMutex.RLock()
	defer r.WMutex.RUnlock()

	return r.W
}

func (r *Rotation) calculateRot(w int) {
	r.WMutex.Lock()
	defer r.WMutex.Unlock()

	r.W = w
}

// ---- HTTP Handler ----
// prefix with 'handle'
func handleGetRotation(w http.ResponseWriter, r *http.Request) {
	// Logic
}

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
