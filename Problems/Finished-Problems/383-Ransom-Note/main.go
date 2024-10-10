package main

import "fmt"

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

func main() {
	ransomNote, magazine := "a", "b"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote, magazine = "aa", "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
	ransomNote, magazine = "aa", "aab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {

	magazineArray, noteArray := []byte(magazine), []byte(ransomNote)

	var m map[byte]int
	m = make(map[byte]int) // stores count of each char

	for _, val := range magazineArray {
		m[val] = m[val] + 1
	}
	for _, val := range noteArray {
		if m[val] == 0 {
			return false // we are out of that character
		} else {
			m[val] = m[val] - 1
		}
	}

	return true
}
