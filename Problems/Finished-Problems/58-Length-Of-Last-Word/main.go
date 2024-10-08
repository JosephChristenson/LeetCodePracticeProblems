package main

import (
	"fmt"
	"slices"
)

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal
// substring
//  consisting of non-space characters only.

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
	s = "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	byteArray := []byte(s)

	slices.Reverse(byteArray)
	result := 0
	wordStart := false
	for _, val := range byteArray {
		if wordStart && val == ' ' {
			return result //a word has ended
		} else if wordStart {
			result++ //a word continues at the end of the sentence
		} else if val != ' ' {
			result++
			wordStart = true //a word has begun at the end of the sentence
		} else {
			//in the middle of at least one space at the end of the sentence. Nothing important to do
		}
	}
	return result // Only one word existed in string

}
