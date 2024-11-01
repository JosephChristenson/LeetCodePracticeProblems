package main

import "fmt"

func main() {
	word1 := "abc"
	word2 := "bca"
	fmt.Println(closeStrings(word1, word2))
}

func closeStrings(word1 string, word2 string) bool {
	word1Byte := []byte(word1)
	word2Byte := []byte(word2)

	if len(word1Byte) != len(word2Byte) {
		return false
	}

	var characterSwap1 byte
	var characterSwap2 byte

	for index := 0; index < len(word2Byte); index++ {
		if word1Byte[index] != word2Byte[index] {
			if characterSwap1 == ' ' {
				characterSwap1 = word1Byte[index]
				characterSwap2 = word2Byte[index]
			} else {
				if word1Byte[index] != characterSwap2 || word2Byte[index] != characterSwap1 {
				}
			}
		}
	}

	return true
}
