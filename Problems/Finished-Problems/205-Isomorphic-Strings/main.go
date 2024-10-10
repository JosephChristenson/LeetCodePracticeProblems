package main

import "fmt"

// Given two strings s and t, determine if they are isomorphic.

// Two strings s and t are isomorphic if the characters in s can be replaced to get t.

// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

func main() {
	s, t := "egg", "add"
	fmt.Println(isIsomorphic(s, t))
	s, t = "foo", "bar"
	fmt.Println(isIsomorphic(s, t))
	s, t = "paper", "title"
	fmt.Println(isIsomorphic(s, t))
	s, t = "badc", "baba"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {

	var m map[byte]byte // used to keep track of what a letter is switched to
	m = make(map[byte]byte)
	var used map[byte]byte // used to see if a character has been used already
	used = make(map[byte]byte)

	arrayS, arrayT := []byte(s), []byte(t)

	test := m['a']              // find value of new map
	for index := range arrayS { // array lengths are the same
		if m[arrayS[index]] != arrayT[index] {
			if m[arrayS[index]] == test && used[arrayT[index]] == test { // if the map hasn't been initialized yet set it to the new value
				m[arrayS[index]] = arrayT[index]
				used[arrayT[index]] = arrayS[index]
			} else {
				return false
			}

		} else {
			//current Char is good
		}
	}
	return true
}
