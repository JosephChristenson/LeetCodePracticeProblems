package main

import (
	"fmt"
)

// The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

// countAndSay(1) = "1"
// countAndSay(n) is the run-length encoding of countAndSay(n - 1).
// Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

// Given a positive integer n, return the nth element of the count-and-say sequence.

func main() {
	n := 4
	fmt.Println(countAndSay(n))
	n = 1
	fmt.Println(countAndSay(n))
}

func countAndSay(n int) string {
	result := "1"
	for n > 1 {
		result = runLengthEncoding(result)
		n--
	}
	return result
}

func runLengthEncoding(n string) string {
	byteArray := []byte(n)
	result := []byte{}
	count := 1
	index := 1
	prev := byteArray[0]
	for index <= len(byteArray)-1 {
		if prev == byteArray[index] {
			count++
		} else {
			result = append(result, byte('0'+count))
			result = append(result, prev)
			prev = byteArray[index]
			count = 1
		}
		index++
	}
	result = append(result, byte('0'+count))
	result = append(result, prev)
	return string(result)
}
