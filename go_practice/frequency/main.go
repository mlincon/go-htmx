/*
In this question, we are building a tool that analyzes the frequency of words
in a given text. You need to implement a function wordFrequency that receives a
string and returns a map with the frequency of each word in the string.
*/

package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	var words []string = strings.Fields(text)

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}
	return frequency
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))
}
