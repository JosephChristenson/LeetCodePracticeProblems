package main

import "fmt"

func main() {
	word1, word2 := "abc", "pqr"
	fmt.Println(mergeAlternately(word1, word2))
	word1, word2 = "ab", "pqrs"
	fmt.Println(mergeAlternately(word1, word2))
	word1, word2 = "abcd", "pq"
	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	byteArray1, byteArray2 := []byte(word1), []byte(word2)
	result := make([]byte, 0, len(byteArray1)+len(byteArray2)-2)

	index1, index2 := 0, 0

	for index1 < len(byteArray1) && index2 < len(byteArray2) {
		result = append(result, byteArray1[index1])
		result = append(result, byteArray2[index2])
		index1++
		index2++
	}
	for index1 < len(byteArray1) {
		result = append(result, byteArray1[index1])
		index1++
	}
	for index2 < len(byteArray2) {
		result = append(result, byteArray2[index2])
		index2++
	}
	return string(result)
}
