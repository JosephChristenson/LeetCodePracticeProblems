package main

import (
	"fmt"
	"strconv"
)

// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func main() {
	x := 123
	fmt.Println(reverse(x))
	x = -123
	fmt.Println(reverse(x))
	x = 120
	fmt.Println(reverse(x))

}

func reverse(x int) int {
	byteArray := []byte(strconv.Itoa(x))
	result := []byte{}
	if len(byteArray) <= 1 {
		return x
	}
	negative := x < 0

	index := len(byteArray) - 1

	for index > 0 {
		result = append(result, byteArray[index])
		index--
	}
	if !negative {
		result = append(result, byteArray[0]) // only take the first char if it isn't the negative sign
	}

	i, _ := strconv.Atoi(string(result))
	if negative {
		i = i * -1
	}

	if i >= 2147483648 || i < -2147483648 {
		return 0
	}
	return i
}
