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
	start := &chart{}
	start = createGrid(start, grid, 0, 0)
	results := make([]string, 0, len(words))
	for _, val := range words {
		word := []byte(val)
		rowhead := start
		for rowhead != nil {
			curr := rowhead
			foundWord := false
			for curr != nil {
				if curr.val == word[0] {
					foundWord = foundWord || searchAround(word, curr, 9)
				}
				curr = curr.directions[2]
			}
			if foundWord {
				results = append(results, val)
				break
			}
			rowhead = rowhead.directions[4]
		}
	}
	return results
}

func searchAround(word []byte, root *chart, direction int) bool {
	if root == nil {
		return false
	}
	if len(word) != 1 {
		switch direction {
		case 9:
			result := false
			for i := 0; i < 8; i++ {
				result = result || searchAround(word[1:], root, i)
			}
			return result
		default:
			return searchAround(word[1:], root.directions[direction], direction)
		}
	} else {
		return root.val == word[0]
	}
}

type chart struct {
	val        byte
	directions [8]*chart
}

func createGrid(root *chart, grid [][]byte, X int, Y int) *chart {
	root.val = grid[Y][X]
	if X < len(grid[0])-1 {
		newChild := &chart{}
		newChild = createGrid(newChild, grid, X+1, Y)
		newChild.directions[6] = root
		root.directions[2] = newChild
	}
	if Y < len(grid)-1 {
		newChild := &chart{}
		newChild = createGrid(newChild, grid, X, Y+1)
		newChild.directions[0] = root
		root.directions[4] = newChild
	}

	if root.directions[0] != nil {
		root.directions[7] = root.directions[0].directions[6]
		root.directions[1] = root.directions[0].directions[2]
		root.directions[0].directions[2] = root.directions[7]
		root.directions[0].directions[6] = root.directions[1]
	}
	if root.directions[4] != nil {
		root.directions[5] = root.directions[4].directions[6]
		root.directions[3] = root.directions[4].directions[2]
		root.directions[4].directions[2] = root.directions[3]
		root.directions[4].directions[6] = root.directions[5]
	}
	return root
}
