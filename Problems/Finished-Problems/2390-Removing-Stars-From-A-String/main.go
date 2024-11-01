package main

import (
	"fmt"
	"slices"
)

// You are given a string s, which contains stars *.

// In one operation, you can:

// Choose a star in s.
// Remove the closest non-star character to its left, as well as remove the star itself.
// Return the string after all stars have been removed.

func main() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func removeStars(s string) string {
	skips := 0
	byteArray := []byte(s)
	results := make([]byte, 0, len(byteArray))
	for index := len(byteArray) - 1; index >= 0; index-- {
		if byteArray[index] == '*' {
			skips++
		} else {
			if skips > 0 {
				skips--
			} else {
				results = append(results, byteArray[index])
			}
		}
	}
	slices.Reverse(results)
	return string(results)
}
