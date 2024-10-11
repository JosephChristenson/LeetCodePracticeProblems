package main

import "fmt"

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
	s = "(]"
	fmt.Println(isValid(s))
	s = "([])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	byteArray := []byte(s)

	if len(byteArray)%2 == 1 {
		return false
	}
	index := 0

	for len(byteArray) > 0 {
		if index+1 == len(byteArray) {
			return false
		}
		if byteArray[index] == '(' && byteArray[index+1] == ')' || byteArray[index] == '[' && byteArray[index+1] == ']' || byteArray[index] == '{' && byteArray[index+1] == '}' {
			if len(byteArray) == 2 {
				return true
			}
			byteArray = append(byteArray[:index], byteArray[(index+2):]...)
			if index < 2 {
				index = 0
			} else {
				index = index - 2 // reset back two spaces
			}
		} else {
			index++
		}
	}
	return true
}
