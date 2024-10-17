package main

import (
	"fmt"
	"math"
)

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.

// The algorithm for myAtoi(string s) is as follows:

// Whitespace: Ignore any leading whitespace (" ").
// Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity is neither present.
// Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
// Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
// Return the integer as the final result.

func main() {
	s := "42"
	fmt.Println(myAtoi(s))
	s = "-042"
	fmt.Println(myAtoi(s))
	s = "1337c0d3"
	fmt.Println(myAtoi(s))
	s = "0-1"
	fmt.Println(myAtoi(s))
	s = "words and 987"
	fmt.Println(myAtoi(s))
	s = "20000000000000000000"
	fmt.Println(myAtoi(s))

}

func myAtoi(s string) int {
	byteArray := []byte(s)

	negative := false
	negativeDetermined := false
	whitespaceOver := false
	breakOut := false

	result := []int{}
	for _, val := range byteArray {
		if len(result) > 0 {
			whitespaceOver = true
			negativeDetermined = true
		}
		switch val {
		case ' ':
			if whitespaceOver {
				breakOut = true
			}
		case '-':
			if negativeDetermined == false {
				negative = true
				negativeDetermined = true
				whitespaceOver = true
			} else {
				breakOut = true
			}
		case '+':
			if negativeDetermined == false {
				negativeDetermined = true
				whitespaceOver = true
			} else {
				breakOut = true
			}
		case '0':
			result = append(result, 0)
		case '1':
			result = append(result, 1)
		case '2':
			result = append(result, 2)
		case '3':
			result = append(result, 3)
		case '4':
			result = append(result, 4)
		case '5':
			result = append(result, 5)
		case '6':
			result = append(result, 6)
		case '7':
			result = append(result, 7)
		case '8':
			result = append(result, 8)
		case '9':
			result = append(result, 9)
		default:
			breakOut = true
		}
		if breakOut {
			break
		}
	}
	float := float64(0)

	for index, val := range result {
		float = float + float64(val)*math.Pow10(len(result)-index-1)
	}

	if negative {
		float = float * -1
	}
	if float >= 2147483647 {
		float = 2147483647
	} else if float <= -2147483648 {
		float = -2147483648
	}
	return int(float)
}
