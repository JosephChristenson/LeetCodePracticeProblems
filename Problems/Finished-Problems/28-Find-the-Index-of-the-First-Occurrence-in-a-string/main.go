package main

import "fmt"

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
	haystack = "leetcode"
	needle = "leeto"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {

	byteHay := []byte(haystack)
	byteNeedle := []byte(needle)

	possibleindexes := []int{}

	if len(byteHay) >= len(byteNeedle) {

		for ind, value := range byteHay {
			if value == byteNeedle[0] {
				possibleindexes = append(possibleindexes, ind)
			}
		}
	}

	for _, value := range possibleindexes {
		if value+len(byteNeedle) <= len(byteHay) {
			if string(byteHay[value:value+len(byteNeedle)]) == needle {
				return value
			}
		}
	}

	return -1

}
