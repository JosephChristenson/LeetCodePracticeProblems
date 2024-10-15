package main

import (
	"fmt"
)

func main() {

	s := "aa"
	p := "a"
	fmt.Println(isMatch(s, p))
	s = "aa"
	p = "a*"
	fmt.Println(isMatch(s, p))
	s = "ab"
	p = ".*"
	fmt.Println(isMatch(s, p))

}
func isMatch(s string, p string) bool {
	root := Trie{root: &trieNode{}} // create new trie
	root.Insert(s)
	return root.Search(p)
}

type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}
type Trie struct {
	root *trieNode
}

func (this *Trie) Insert(word string) {
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

func (this *Trie) Search(word string) bool {
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
		} else if word[i] == '*' {
			if 

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
	return false
}
