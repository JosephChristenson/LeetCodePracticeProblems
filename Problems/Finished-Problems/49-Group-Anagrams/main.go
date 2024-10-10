package main

import (
	"fmt"
	"sort"
)

// Given an array of strings strs, group the
// anagrams
//  together. You can return the answer in any order.

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	strs = []string{""}
	fmt.Println(groupAnagrams(strs))
	strs = []string{"a"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}

	var m map[string][]string
	m = make(map[string][]string)

	for _, val := range strs {
		test := []byte(val)
		sort.Slice(test, func(i, j int) bool {
			return test[i] < test[j]
		})

		m[string(test)] = append(m[string(test)], val)
	}
	result := [][]string{}
	for _, array := range m {
		result = append(result, array)
	}
	return result
}
