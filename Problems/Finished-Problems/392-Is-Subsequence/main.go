package main

import "fmt"

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	searchFor := []byte(s)
	searchThis := []byte(t)
	wordIndex := 0
	searchIndex := 0

	for searchIndex < len(searchThis) && wordIndex < len(searchFor) {
		if searchFor[wordIndex] == searchThis[searchIndex] {
			wordIndex++
		}
		searchIndex++
	}

	return wordIndex == len(searchFor)
}
