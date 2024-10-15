package main

import (
	"fmt"
)

func main() {

	commands := []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"}
	num := [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}}
	runFunc(commands, num)

}

func runFunc(input []string, values [][]string) {
	var trie WordDictionary

	for i, s := range input {
		if s == "WordDictionary" {
			trie = Constructor()
		} else if s == "addWord" {
			trie.AddWord(values[i][0])
		} else if s == "search" {
			fmt.Println(trie.Search(values[i][0]))
		}
	}
}

type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}
type WordDictionary struct {
	root *trieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &trieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	wordLength := len(word)
	curr := this.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if curr.childrens[index] == nil { // if the next letter of the tree doesn't already exist make it
			curr.childrens[index] = &trieNode{}
		}
		curr = curr.childrens[index] // set the branch to the next letter
	}
	curr.isWordEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	wordLength := len(word)
	curr := this.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if word[i] == '.' {
			if i+1 == wordLength { // end of word already shortcut
				for _, possibleChild := range curr.childrens {
					if possibleChild != nil {
						if possibleChild.isWordEnd {
							return true
						}
					}
				}
				return false
			} else {
				search := false
				for j := 0; j < 26; j++ {
					if curr.childrens[j] != nil {
						search = searchAll(curr.childrens[j], string(word[i+1:]))
						if search {
							return true
						}
					}
				}
				return false
			}
		} else {
			if curr.childrens[index] == nil {
				return false // we have reached the end of a branch without getting the full word
			}
			curr = curr.childrens[index] // only increment if . not reached
		}
	}
	return curr.isWordEnd

}

func searchAll(node *trieNode, word string) bool {
	wordLength := len(word)
	curr := node

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if word[i] == '.' {
			if i+1 == wordLength { // end of word already shortcut
				for _, possibleChild := range curr.childrens {
					if possibleChild != nil {
						if possibleChild.isWordEnd {
							return true
						}
					}
				}
				return false
			} else {
				search := false
				for j := 0; j < 26; j++ {
					if curr.childrens[j] != nil {
						search = searchAll(curr.childrens[j], string(word[i+1:]))
						if search {
							return true
						}
					}
				}
				return false
			}
		} else {
			if curr.childrens[index] == nil {
				return false // we have reached the end of a branch without getting the full word
			}
			curr = curr.childrens[index] // only increment if . not reached
		}
	}
	return curr.isWordEnd
}
