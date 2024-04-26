package crypt

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Takes in a string and performs a letter-based frequency analysis
// and prints out in descending order
func FreqAnalysis(content string) {

	freqMap := make(map[string]int)
	matcher, _ := regexp.CompilePOSIX("[a-zA-Z]")

	for pos, char := range content {
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
