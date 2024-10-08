package main

// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.

// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{3, 0, 6, 1, 5}
	fmt.Println(hIndex(citations))
	citations = []int{1, 3, 1}
	fmt.Println(hIndex(citations))

}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	for len(citations) > 0 {
		if citations[len(citations)-1] >= len(citations) {
			return len(citations)
		} else {
			citations = citations[:len(citations)-1]
		}
	}

	return 0
}
