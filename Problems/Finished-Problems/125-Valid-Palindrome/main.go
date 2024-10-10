package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s = "race a car"
	fmt.Println(isPalindrome(s))
	s = " "
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]")
	onlyLetters := reg.ReplaceAllString(s, "")

	return checkPalindrome(strings.ToLower(onlyLetters))

}

func checkPalindrome(x string) bool {
	byteArray := []byte(x)
	for i, _ := range byteArray {
		if byteArray[i] != byteArray[len(byteArray)-(i+1)] {
			return false
		}
	}
	return true
}
