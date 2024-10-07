package main

import "fmt"

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

func main() {
	s := "III"
	fmt.Print(romanToInt(s))
	s = "LVIII"
	fmt.Print(romanToInt(s))
	s = "MCMXCIV"
	fmt.Print(romanToInt(s))
}

func romanToInt(s string) int {
	result := 0
	byteArray := []byte(s)
	previousLetter := byte('a')

	for _, Letter := range byteArray {
		switch Letter {
		case 'I':
			result = result + 1
			previousLetter = 'I'
		case 'V':
			if previousLetter == 'I' {
				result = result + 3
			} else {
				result = result + 5
			}
			previousLetter = 'V'
		case 'X':
			if previousLetter == 'I' {
				result = result + 8
			} else {
				result = result + 10
			}
			previousLetter = 'X'
		case 'L':
			if previousLetter == 'X' {
				result = result + 30
			} else {
				result = result + 50
			}
			previousLetter = 'L'
		case 'C':
			if previousLetter == 'X' {
				result = result + 80
			} else {
				result = result + 100
			}
			previousLetter = 'C'
		case 'D':
			if previousLetter == 'C' {
				result = result + 300
			} else {
				result = result + 500
			}
			previousLetter = 'D'
		case 'M':
			if previousLetter == 'C' {
				result = result + 800
			} else {
				result = result + 1000
			}
			previousLetter = 'L'
		default:
		}
	}

	return result
}
