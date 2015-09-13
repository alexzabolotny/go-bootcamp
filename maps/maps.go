package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

// WordCount counting words
func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, n := range strings.Fields(s) {
		result[n]++
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
