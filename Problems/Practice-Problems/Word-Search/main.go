package main

import "fmt"

func main() {
	wordSearch := [][]byte{
		{'b', 'a', 'n', 'a', 'n'},
		{'j', 'a', 'c', 'k', 'f'},
		{'c', 'a', 'n', 'a', 'r'},
		{'v', 'e', 'g', 'e', 't'},
		{'o', 'r', 'a', 'n', 'g'},
	}
	stringsToFind := []string{
		"cat",
		"nfr",
		"aec",
		"ovcjb",
		"gnaro"}

	fmt.Println(WordSearchFunc(wordSearch, stringsToFind))
}

func WordSearchFunc(grid [][]byte, words []string) []string {
	results := make([]string, 0, len(words))
	for _, word := range words {
		if wordFound(grid, word) {
			results = append(results, word)
		}
	}
	return results
}

func wordFound(grid [][]byte, word string) bool {
	byteArray := []byte(word)
	lengthWord := len(byteArray)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byteArray[0] {
				checks := [8]bool{false, false, false, false, false, false, false, false} // Up, UpR, R, DownR, Down,DownL,L,UpL
				if i+1 >= lengthWord {
					checks[0] = true
					//search up
				}
				if len(grid)-i >= lengthWord {
					checks[4] = true
					// search down
				}
				if j+1 >= lengthWord {
					checks[6] = true
					if checks[0] {
						checks[7] = true
					}
					if checks[4] {
						checks[5] = true
					}
					// search left
				}
				if len(grid[0])-j >= lengthWord {
					checks[2] = true
					if checks[0] {
						checks[1] = true
					}
					if checks[4] {
						checks[3] = true
					}
					// search right
				}
				index := 1
				for index < lengthWord && (checks[0] || checks[1] || checks[2] || checks[3] || checks[4] || checks[5] || checks[6] || checks[7]) {
					if checks[0] {
						checks[0] = grid[i-index][j] == byteArray[index]
					}
					if checks[1] {
						checks[1] = grid[i-index][j+index] == byteArray[index]
					}
					if checks[2] {
						checks[2] = grid[i][j+index] == byteArray[index]
					}
					if checks[3] {
						checks[3] = grid[i+index][j+index] == byteArray[index]
					}
					if checks[4] {
						checks[4] = grid[i+index][j] == byteArray[index]
					}
					if checks[5] {
						checks[5] = grid[i+index][j-index] == byteArray[index]
					}
					if checks[6] {
						checks[6] = grid[i][j-index] == byteArray[index]
					}
					if checks[7] {
						checks[7] = grid[i-index][j-index] == byteArray[index]
					}
					index++
				}
				if index == lengthWord && (checks[0] || checks[1] || checks[2] || checks[3] || checks[4] || checks[5] || checks[6] || checks[7]) {
					fmt.Println(i, j)
					return true
				}
			}
		}
	}
	return false
}
