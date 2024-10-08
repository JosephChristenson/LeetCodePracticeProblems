package main

import "fmt"

func main() {
	digits := string("23")
	fmt.Println(letterCombinations(digits))
	digits = string("")
	fmt.Println(letterCombinations(digits))
	digits = string("2")
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	byteArray := []byte(digits)

	result := make([]string, 0, 256)
	var m map[byte][]byte
	m = make(map[byte][]byte)

	m['2'] = []byte{'a', 'b', 'c'}
	m['3'] = []byte{'d', 'e', 'f'}
	m['4'] = []byte{'g', 'h', 'i'}
	m['5'] = []byte{'j', 'k', 'l'}
	m['6'] = []byte{'m', 'n', 'o'}
	m['7'] = []byte{'p', 'q', 'r', 's'}
	m['8'] = []byte{'t', 'u', 'v'}
	m['9'] = []byte{'w', 'x', 'y', 'z'}

	for index, val := range byteArray {
		if index == 0 {
			for _, value := range m[val] {
				result = append(result, string(value))
			}
		} else {
			result = allCombos(result, m[val])
		}
	}
	return result
}

func allCombos(a []string, b []byte) []string {
	result := make([]string, 0, 256)
	for _, valA := range a {
		for _, valB := range b {
			result = append(result, valA+string(valB))
		}
	}
	return result
}
