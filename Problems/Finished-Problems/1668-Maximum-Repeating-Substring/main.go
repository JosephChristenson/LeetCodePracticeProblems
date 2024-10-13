package main

import "fmt"

func main() {
	sequence, word := "ababc", "ab"
	// fmt.Println(maxRepeating(sequence, word))
	// sequence, word = "ababc", "ba"
	// fmt.Println(maxRepeating(sequence, word))
	// sequence, word = "ababc", "ac"
	// fmt.Println(maxRepeating(sequence, word))
	sequence, word = "baba", "b"
	fmt.Println(maxRepeating(sequence, word))
	sequence, word = "aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba"
	fmt.Println(maxRepeating(sequence, word))
}
func maxRepeating(sequence string, word string) int {
	if sequence == word {
		return 1
	}
	byteSeq, byteWord := []byte(sequence), []byte(word)

	result, index := 0, 0
	repeating := false
	max := 0

	for index+len(byteWord) <= len(byteSeq) {
		if byteSeq[index] == byteWord[0] {
			found := true
			for ind, val := range byteWord {
				if byteSeq[index+ind] != val {
					found = false
					break
				}
			}
			if found {
				if len(byteWord) != 1 {
					index = index + len(byteWord) - 1
				}
				if repeating {
					result++
					if max < result {
						max = result
					}
				} else {
					repeating = true
					result = 1
					if max <= result {
						max = result
					}
				}
			} else {
				if repeating {
					repeating = false
					index = index - len(byteWord)
				}
				result = 0
			}
		} else {
			repeating = false
		}
		index++
	}
	return max
}
