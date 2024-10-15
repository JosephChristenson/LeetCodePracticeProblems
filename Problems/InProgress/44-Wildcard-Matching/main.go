package main

import (
	"fmt"
)

// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).

func main() {
	s, p := "aa", "a"
	fmt.Println(isMatch(s, p))
	s, p = "aa", "*"
	fmt.Println(isMatch(s, p))
	s, p = "cb", "?a"
	fmt.Println(isMatch(s, p))

}
func isMatch(s string, p string) bool {
	if p == "*" {
		return true
	}
	stringArray := []byte(s)
	patternArray := []byte(p)

	for index := 0; index <= len(stringArray)-1 && index <= len(patternArray)-1; index++ {
		if patternArray[index] == '*' {
			return true
		} else if patternArray[index] == stringArray[index] {
			// do nothing
		} else if patternArray[index] == '?' {
			// do nothing
		} else {
			return false
		}
	}
	if len(stringArray) == len(patternArray) { // if it made it this far and they are the same lenght then they are good
		return true
	} else {
		return false
	}
}
