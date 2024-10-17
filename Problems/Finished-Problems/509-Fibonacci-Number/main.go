package main

import "fmt"

func main() {
	n := 2
	fmt.Println(fib(n))
	n = 3
	fmt.Println(fib(n))
	n = 5
	fmt.Println(fib(n))
}

func fib(n int) int {
	temp1, temp2 := 1, 0

	if n == 1 {
		return temp1
	}
	if n == 0 {
		return temp2
	}
	index := 2

	for index <= n {
		temp1, temp2 = temp1+temp2, temp1
		index++
	}
	return temp1
}
