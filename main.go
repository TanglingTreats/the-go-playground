package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"math"
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

	mode := flag.String("mode", "default", "The mode to run the program in")
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
	default:
		ASN1Encoding()
	}

}

func cryptoFun(filename string) {
	// Open supplied filename
	file := iofile.ReadFile(filename)

	crypt.FreqAnalysis(string(file))
}

type ASN1_Types int

const (
	BOOLEAN ASN1_Types = iota
	INTEGER
	BIT_STR
	OCTET_STR
	NULL
	OBJ_IDENT
	OBJ_DESC
	EXTERNAL
	REAL
	ENUM
	EMBED_PDV
	UTF8_STR
	CHOICE
	RELATIVE_OID
	TIME
	SEQ
	SEQ_OF
	SET
	SET_OF
)

func ASN1Encoding() {

	var buf bytes.Buffer

	const size = 256

	// egData := make([]byte, size)
	// _, err := rand.Read(egData)

	// if err != nil {
	// fmt.Println("Unable to generate random data")
	// }

	// curSize := encodeASN1(egData, OCTET_STR, &buf)
	intArr := []int{2, 4}
	curSize := 0
	for _, data := range intArr {
		intBuf := new(bytes.Buffer)
		writeAsByteSlice(intBuf, data)
		curSize += encodeASN1(intBuf.Bytes(), INTEGER, &buf)
	}

	// Create new slice and copy data over
	data := make([]byte, curSize)
	copy(data, buf.Bytes())
	buf.Reset() // Reset existing buffer

	curSize = encodeASN1(data, SEQ, &buf)

	fmt.Printf("Current size: %d\n", curSize)
	fmt.Printf("%#v\n", buf.Bytes())
}

/**
* Encode an ASN1 type
 */
func encodeASN1(data []byte, asnType ASN1_Types, buf *bytes.Buffer) int {
	size := 0

	dataLen := len(data)

	switch asnType {
	case INTEGER:
		fmt.Println("Encoding INT")

		_, err := buf.Write([]byte{0x02, byte(dataLen)})

		if err != nil {
			fmt.Println("Failed to write to buffer")
		}
		_, err = buf.Write(data)

		if err != nil {
			fmt.Println("Failed to write to buffer")
		}

		size += 2 + dataLen // INT Tag, Length, Data

	case OCTET_STR:
		if dataLen <= 0x7F {
			fmt.Println("Encoding OCTET STR")
			buf.Write([]byte{0x04, byte(dataLen)})
			buf.Write(data)

			size += 2 + dataLen

		} else {

			fmt.Println("Encoding extended length OCTET STR")

			numOfBytes := getNumOfBytesFromInt(dataLen)

			buf.Write([]byte{0x04, 0x80 | byte(numOfBytes)})
			writeAsByteSlice(buf, dataLen)
			buf.Write(data)

			size += 2 + numOfBytes + dataLen // Tag, Size of Len, Len, Len of Data
		}

	case SEQ:
		fmt.Println("Encoding a SEQUENCE")
		buf.Write([]byte{0x30, byte(dataLen)})
		buf.Write(data)

		size += 2 + dataLen
	case SET:
		fmt.Println("Encoding a SET")
		buf.Write([]byte{0x31, byte(dataLen)})
		buf.Write(data)

		size += 2 + dataLen
	default:
		fmt.Println("No such ASN type handled")
	}

	return size
}

func writeAsByteSlice(buf *bytes.Buffer, i int) {

	switch {
	case i >= math.MinInt8 && i <= math.MaxInt8:
		binary.Write(buf, binary.BigEndian, int8(i))
	case i >= math.MinInt16 && i <= math.MaxInt16:
		binary.Write(buf, binary.BigEndian, int16(i))
	case i >= math.MinInt32 && i <= math.MaxInt32:
		binary.Write(buf, binary.BigEndian, int32(i))
	case i >= math.MinInt64 && i <= math.MaxInt64:
		binary.Write(buf, binary.BigEndian, int64(i))
	default:
		fmt.Println("Integer is out of range")
	}
}

func getNumOfBytesFromInt(i int) int {
	counter := 0

	for i > 0 {
		i >>= 8
		counter += 1
	}

	return counter
}
