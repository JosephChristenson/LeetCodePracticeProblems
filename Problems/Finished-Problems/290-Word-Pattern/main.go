package main

import (
	"fmt"
	"strings"
)

// Given a pattern and a string s, find if s follows the same pattern.

// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:

// Each letter in pattern maps to exactly one unique word in s.
// Each unique word in s maps to exactly one letter in pattern.
// No two letters map to the same word, and no two words map to the same letter.

func main() {
	input, s := "abba", "dog cat cat dog"
	fmt.Println(wordPattern(input, s))
	input, s = "abba", "dog cat cat fish"
	fmt.Println(wordPattern(input, s))
	input, s = "aaaa", "dog cat cat dog"
	fmt.Println(wordPattern(input, s))
}

func wordPattern(pattern string, s string) bool {

	var findChar map[string]byte // used to keep track of what a letter is switched to
	findChar = make(map[string]byte)
	var findWord map[byte]string // used to see if a word has been used
	findWord = make(map[byte]string)

	arrayChar := []byte(pattern)
	arrayWord := strings.Split(s, " ")

	if len(arrayChar) != len(arrayWord) {
		return false
	}
	testChar := findWord['a'] // find value of new map
	testWord := findChar["a"]

	for index := range arrayChar { // array lengths are the same
		if !(findChar[arrayWord[index]] == arrayChar[index]) || !(findWord[arrayChar[index]] == arrayWord[index]) { // check if char and word are already matched
			if findChar[arrayWord[index]] == testWord { // word hasn't been asigned yet
				findChar[arrayWord[index]] = arrayChar[index]
			} else { // already been asigned
				return false
			}
			if findWord[arrayChar[index]] == testChar { // word hasn't been asigned yet
				findWord[arrayChar[index]] = arrayWord[index]
			} else { // already been asigned
				return false
			}
		} else {
			//current Char is good
		}
	}
	return true
}
