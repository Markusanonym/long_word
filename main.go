package main

import (
	"fmt"
	"regexp"
	"strings"
)

func LongestWord(sen string) string {
	re := regexp.MustCompile(`[^\w\s]`) // Remove punctuation
	sen = re.ReplaceAllString(sen, "")

	words := strings.Split(sen, " ") // Split the string into words
	longestWord := ""

	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return longestWord
}

func main() {
	fmt.Println(LongestWord("Hello, World!")) // Output: "Hello"
}
