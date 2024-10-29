package main

import "fmt"

// Given a string s, reverse only all the vowels in the string and return it.

// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

func main() {
	s := "IceCreAm"
	fmt.Println(reverseVowels(s))
	s = "leetcode"
	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	byteArray := []byte(s)
	left, right := 0, len(byteArray)-1

	for left <= right {
		if isVowel(byteArray[left]) && isVowel(byteArray[right]) {
			byteArray[left], byteArray[right] = byteArray[right], byteArray[left]
			left++
			right--
		} else {
			if !isVowel(byteArray[left]) {
				left++
			}
			if !isVowel(byteArray[right]) {
				right--
			}
		}
	}

	return string(byteArray)
}

func isVowel(char byte) bool {
	switch char {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	case 'A':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'O':
		return true
	case 'U':
		return true
	default:
		return false
	}
}
