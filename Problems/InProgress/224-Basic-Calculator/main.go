package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1 + 1"
	// fmt.Println(calculate(s))
	// s = " 2-1 + 2 "
	// fmt.Println(calculate(s))
	s = "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	byteArray := []byte(strings.ReplaceAll(s, " ", ""))
	result, _ := calculateParen(byteArray)
	return result
}

func calculateParen(input []byte) (int, int) { // return total, return index

	prev, curr := 0, 0
	foundFirst := false
	opperation := ' '

	// num := []byte{} // used for number converting
	for index, val := range input {
		switch val {
		case '(':
			if foundFirst {
				curr, index = calculateParen(input[index+1:])
				prev = returnValue(prev, curr, byte(opperation))
				curr = 0
			} else {
				foundFirst = true
				prev, _ = calculateParen(input[index+1:])
				if opperation == '-' {
					prev = prev * -1
				}
			}
		case ')':
			return prev, index + 1
		case '+':
			opperation = '+'
		case '-':
			opperation = '-'
		default:
			if !foundFirst {
				foundFirst = true
				prev, _ = strconv.Atoi(string(val))
				if opperation == '-' {
					prev = prev * -1
				}
			} else {
				curr, _ = strconv.Atoi(string(val))
				prev = returnValue(prev, curr, byte(opperation))
				curr = 0
			}
		}
	}
	return prev, 0
}

func returnValue(prev int, curr int, opp byte) int {
	result := 0
	switch opp {
	case '-':
		result = prev - curr
	case '+':
		result = prev + curr
	default:
		fmt.Println("Awwoogaaaa")
	}
	return result
}
