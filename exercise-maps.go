package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int = make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		m[word] += 1
	}
	return m
}

func main() {
	fmt.Println(WordCount("The quick brown fox jumps over the lazy dog"))
}
