package main

import (
	"fmt"
)

// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
	str1 = "ABABAB"
	str2 = "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))
	str1 = "LEET"
	str2 = "CODE"
	fmt.Println(gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	byteArray1 := []byte(str1)
	byteArray2 := []byte(str2)

	var smallestLength int
	if len(byteArray1) < len(byteArray2) {
		smallestLength = len(byteArray1)
	} else {
		smallestLength = len(byteArray2)
	}
	length := smallestLength
	divisor := 1

	for smallestLength >= divisor {
		if len(byteArray1)%(length) == 0 && len(byteArray2)%(length) == 0 {
			pattern := string(byteArray1[:length])
			patternGood := true
			for integer := 0; integer < len(byteArray1); integer += length {
				if pattern != string(byteArray1[integer:integer+length]) {
					patternGood = false
					break
				}
			}
			if patternGood {
				for integer := 0; integer < len(byteArray2); integer += length {
					if pattern != string(byteArray2[integer:integer+length]) {
						patternGood = false
						break
					}
				}
			}
			if patternGood {
				return pattern
			}

		}
		divisor++
		length = smallestLength / divisor
	}
	return ""
}
