package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))
	arr = []int{1, 2}
	fmt.Println(uniqueOccurrences(arr))
	arr = []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))

}

func uniqueOccurrences(arr []int) bool {
	var count map[int]int
	count = make(map[int]int)
	for _, val := range arr {
		count[val]++
	}

	var exist map[int]bool
	exist = make(map[int]bool)

	for _, val := range count {
		if exist[val] {
			return false
		} else {
			exist[val] = true
		}
	}
	return true
}
