package main

import (
	"fmt"
	"strconv"
)

// Given an array of characters chars, compress it using the following algorithm:

// Begin with an empty string s. For each group of consecutive repeating characters in chars:

// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

// After you are done modifying the input array, return the new length of the array.

// You must write an algorithm that uses only constant extra space.

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
	chars = []byte{'a'}
	fmt.Println(compress(chars))
	chars = []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars))
}

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}
	var previous byte
	count := 1
	countindex := 0 //stores where the count should be displayed

	for index, val := range chars {
		if index == 0 {
			previous = val
		} else if val == previous {
			if count == 1 {
				countindex = index
			}
			chars[index] = ' '
			count++
		} else {
			previous = val
			if count > 1 {
				countreport := []byte(strconv.Itoa(count))
				for dis, val := range countreport {
					chars[countindex+dis] = val
				}
				count = 1
				countindex = index + 1
			}
		}
	}

	if count > 1 {
		countreport := []byte(strconv.Itoa(count))
		for dis, val := range countreport {
			chars[countindex+dis] = val
		}
	}
	replaceAt := 0
	for index, val := range chars {
		if val != ' ' {
			char := val
			chars[index] = ' '
			chars[replaceAt] = char
			replaceAt++
		}
	}

	return replaceAt
}
