package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
	strs = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	result, test := "", ""
	for _, char := range []byte(strs[0]) {
		test += string(char)
		for _, value := range strs {
			if !strings.HasPrefix(value, test) {
				return result
			}
		}
		result = test
	}
	return result
}
