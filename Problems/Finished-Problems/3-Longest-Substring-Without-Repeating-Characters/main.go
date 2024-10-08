package main

import (
	"fmt"
	"slices"
)

// Given a string s, find the length of the longest
// substring
//  without repeating characters.

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	byteArray := []byte(s)
	substring := make([]byte, 0, 26)
	result := 0
	for _, value := range byteArray {
		if !slices.Contains(substring, value) {
			substring = append(substring, value)
			if len(substring) > result {
				result = len(substring)
			}
		} else {
			substring = removeAllByteBeforeFirstRepeat(substring, value)
			substring = append(substring, value)
		}
	}
	return result
}
func removeAllByteBeforeFirstRepeat(bytearray []byte, val byte) []byte {
	newsubstring := make([]byte, 0, 26)
	for index, value := range bytearray {
		if value == val {
			newsubstring = append(newsubstring, bytearray[index+1:]...)
			break
		}
	}
	return newsubstring
}
