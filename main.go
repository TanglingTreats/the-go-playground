package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
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
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	freqMap := make(map[string]int)

	matcher, _ := regexp.CompilePOSIX("[a-zA-Z]")
	for pos, char := range string(file) {
		_ = pos
		token := strings.ToLower(string(char))
		isChar := matcher.MatchString(token)

		if isChar {
			freqMap[token] += 1
		}
	}

	keys := make([]string, 0, len(freqMap))
	for k := range freqMap {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, freqMap[k])
	}
}
