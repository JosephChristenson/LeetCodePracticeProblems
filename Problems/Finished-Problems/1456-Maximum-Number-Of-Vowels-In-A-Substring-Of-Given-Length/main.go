package main

import "fmt"

func main() {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))
	s = "aeiou"
	k = 2
	fmt.Println(maxVowels(s, k))
	s = "leetcode"
	k = 3
	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {
	byteArray := []byte(s)

	count := 0

	for index := 0; index < k; index++ {
		switch byteArray[index] {
		case 'a', 'e', 'i', 'o', 'u': // gives us a base count for amount of vowels
			count++
		default:
		}
	}
	result := count

	for index := 0; index+k < len(byteArray); index++ {
		switch byteArray[index+k] { // if we added a vowel increment
		case 'a', 'e', 'i', 'o', 'u':
			count++
		default:
		}
		switch byteArray[index] { // if we removed a vowel decrement
		case 'a', 'e', 'i', 'o', 'u':
			count--
		default:
		}
		if count > result {
			result = count
		}
	}
	return result
}
