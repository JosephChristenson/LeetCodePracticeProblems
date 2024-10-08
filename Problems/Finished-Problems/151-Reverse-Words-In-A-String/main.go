package main

import (
	"fmt"
	"slices"
	"strings"
)

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal
// substring
//  consisting of non-space characters only.

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
	s = "  hello world  "
	fmt.Println(reverseWords(s))
	s = "a good   example"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	byteArray := []byte(s)
	currentWord := make([]byte, 0, 1000) //made this big in case the entire string is a word
	result := []string{}

	wordStart := false
	for _, val := range byteArray {
		if wordStart && val == ' ' { //Word ended
			result = append(result, string(currentWord))
			currentWord = make([]byte, 0, 1000)
			wordStart = false
		} else if wordStart { // add letter
			currentWord = append(currentWord, val)
		} else if val != ' ' {
			wordStart = true
			currentWord = append(currentWord, val)
		}
	}
	if string(currentWord) != "" {
		result = append(result, string(currentWord)) // add the last word
	}
	slices.Reverse(result) // put them in reverse
	return strings.Join(result, " ")
}
