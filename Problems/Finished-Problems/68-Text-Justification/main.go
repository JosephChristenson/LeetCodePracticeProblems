package main

import (
	"fmt"
	"strings"
)

// Given an array of strings words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

// You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

// Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line does not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

// For the last line of text, it should be left-justified, and no extra space is inserted between words.

// Note:

// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	fmt.Println(fullJustify(words, maxWidth))
	words = []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth = 16
	fmt.Println(fullJustify(words, maxWidth))
	words = []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth = 20
	fmt.Println(fullJustify(words, maxWidth))

}

func fullJustify(words []string, maxWidth int) []string {
	charactersneeded, wordcount, index := 0, 0, 0

	result := []string{}

	for index < len(words) {
		if wordcount == 0 {
			wordcount++
			charactersneeded = len([]byte(words[index]))
		} else if charactersneeded+1+len([]byte(words[index])) <= maxWidth {
			wordcount++
			charactersneeded = charactersneeded + 1 + len([]byte(words[index]))
		} else {
			index--
			// extraspaces is the amount of spaces needed. we already account for 1 space per word so need to fix that
			// extraspaces = the width of the line minus the characters already used + amount of spaces between words
			extraspaces := maxWidth - charactersneeded + wordcount - 1

			spacesPerWord := extraspaces
			remainder := 0

			if wordcount != 1 {
				spacesPerWord = extraspaces / (wordcount - 1)
				remainder = extraspaces % (wordcount - 1)
			}
			wordsToAdd := words[(index - wordcount + 1) : index+1]
			line := []byte{}

			for index, val := range wordsToAdd {
				for _, char := range []byte(val) {
					line = append(line, char)
				}
				if index < wordcount-1 {
					if remainder > 0 {
						line = append(line, ' ')
						remainder--
					}
					i := 0
					for i < spacesPerWord {
						i++
						line = append(line, ' ')
					}
				}
			}

			for len(line) < maxWidth {
				line = append(line, ' ')
			}

			result = append(result, string(line))
			wordcount = 0
			charactersneeded = 0
		}
		index++
	}
	if wordcount > 0 { //extra words need to be added for the last line
		value := strings.Join(words[len(words)-wordcount:], " ")
		val := []byte(value)
		for len(val) < maxWidth {
			val = append(val, ' ')
		}
		result = append(result, string(val))
	}
	return result
}
