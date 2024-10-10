package main

import "fmt"

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))
	s, t = "rat", "car"
	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	byteS, byteT := []byte(s), []byte(t)
	if len(byteS) != len(byteT) {
		return false
	}

	var m map[byte]int
	m = make(map[byte]int)

	for _, val := range byteS {
		m[val] = m[val] + 1
	}
	for _, val := range byteT {
		if m[val] == 0 {
			return false
		} else {
			m[val] = m[val] - 1
		}
	}
	return true
}
