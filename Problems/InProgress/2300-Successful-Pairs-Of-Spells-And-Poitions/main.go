package main

import (
	"fmt"
	"sort"
)

func main() {
	spells := []int{5, 1, 3}
	potions := []int{1, 2, 3, 4, 5}
	success := 7
	fmt.Println(successfulPairs(spells, potions, int64(success)))
	spells = []int{3, 1, 2}
	potions = []int{8, 5, 8}
	success = 16
	fmt.Println(successfulPairs(spells, potions, int64(success)))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	results := []int{}
	for ind, val := range spells {
		for index := 1; index <= len(potions); index++ {
			if int64(potions[index-1]*val) >= success {
				results = append(results, len(potions)-(index-1))
				break
			}
		}
		if ind+1 > len(results) {
			results = append(results, 0)
		}
	}
	return results
}
