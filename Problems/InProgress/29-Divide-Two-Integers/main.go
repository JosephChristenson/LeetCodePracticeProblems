package main

import "fmt"

func main() {
	dividend := 10
	divisor := 3
	fmt.Println(divide(dividend, divisor))
	dividend = 7
	divisor = -3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	negativeresult := false
	result := 0

	if dividend < 0 {
		dividend = dividend * -1
		negativeresult = !negativeresult
	}
	if divisor < 0 {
		divisor = divisor * -1
		negativeresult = !negativeresult
	}

	if divisor == 1 {
		result = dividend
	} else {

		for dividend >= 0 {
			dividend = dividend - divisor
			if dividend >= 0 {
				result++
			}
		}
	}
	if negativeresult {
		result = result * -1
	}

	//keep result within data limitations
	if result > 2147483647 {
		result = 2147483647
	} else if result < -2147483648 {
		result = 2147483648
	}

	return result
}
