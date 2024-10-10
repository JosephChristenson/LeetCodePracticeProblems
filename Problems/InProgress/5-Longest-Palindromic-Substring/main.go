package main

import "fmt"

// Given a string s, return the longest
// palindromic

// substring
//  in s.

func main(){
	s :="babad"
	fmt.Println(longestPalindrome(s))
	s ="cbbd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	longest := []byte{}
	input := []byte(s)

	if len(input) <= 2{ // make sure the result
		return s
	}

	for center, val{ // given a center point try assuming its by itself the center and iterate off it 
		
		if center == 0 {
			if input[center + 1] == val{ // the first two characters are the same. set the result as both
				longest = val + val
			}else{
				longest = val
			}
		}
		increment := 0

		if input[center - 1]
		for center - increment = 0 && center + increment < len(input){ // stay within the bounds of the string
			
		}

	}
    
}