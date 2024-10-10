package main

import "fmt"

// Given a string s, return the longest
// palindromic

// substring
//  in s.

func main() {
	s := "babad"
	// fmt.Println(longestPalindrome(s))
	// s = "cbbd"
	// fmt.Println(longestPalindrome(s))
	// s = "ac"
	// fmt.Println(longestPalindrome(s))
	// s = "bb"
	// fmt.Println(longestPalindrome(s))
	s = "xaabacxcabaaxcabaax"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	longest := []byte{}
	input := []byte(s)

	if len(input) <= 1 { // make sure the result
		return s
	}

	for center := range input { // given a center point try assuming its by itself the center and iterate off it

		if center == 0 {
			if input[center+1] == input[center] { // the first two characters are the same. set the result as both
				longest = []byte{input[0], input[1]}
			} else {
				longest = []byte{input[0]}
			}
		} else {
			increment := 0
			palindrome := false
			even := 0 // keep track of if the center is 1 number or 2 aba or abba
			odd := false
			if center+1 < len(input) && center-1 >= 0 { // are we at the end?
				if input[center-1] == input[center+1] { // are both our neighbors the same? odd Palindrome
					curr := input[center-increment : center+increment+1]
					if len(curr) > len(longest) {
						longest = curr
					}
					palindrome = true
					odd = true
				}
				if input[center] == input[center+1] { // is it the same as its neighbor to the right? we have a palindrome and need to increment by 1
					curr := input[center : center+2]
					if len(curr) > len(longest) {
						longest = curr
					}
					palindrome = true
					even = 1
				}
			}

			if palindrome {
				for center-increment >= 0 && center+increment < len(input) {
					if odd {
						if input[center-increment] == input[center+increment] { // check odd palindrome
							curr := input[center-increment : center+increment+1]
							if len(curr) > len(longest) {
								longest = curr
							}
						} else {
							odd = false
						}
					}
					if even == 1 {
						if center+increment+even == len(input) {
							break
						}
						if input[center-increment] == input[center+increment+even] { // check even palindrome
							curr := input[center-increment : center+increment+even+1]
							if len(curr) > len(longest) {
								longest = curr
							}
						} else { // palindrome is over
							even = 0
						}
					}
					increment++
				}
			}
		}
	}
	return string(longest)

}
