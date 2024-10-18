package main

import "fmt"

func main() {
	customers := "YYNY"
	fmt.Println(bestClosingTime(customers))
	customers = "NNNNN"
	fmt.Println(bestClosingTime(customers))
	customers = "YYYY"
	fmt.Println(bestClosingTime(customers))
}

func bestClosingTime(customers string) int {

	time := []byte(customers)

	count := 0
	for _, val := range time {
		if val == 'Y' {
			count++
		}
	}
	index := 0

	result := count
	bestIndex := 0
	for index < len(time) {
		if time[index] == 'Y' {
			count = count - 1
		} else {
			count = count + 1
		}
		index++
		if count < result {
			result = count
			bestIndex = index
		}
	}
	return bestIndex
}
