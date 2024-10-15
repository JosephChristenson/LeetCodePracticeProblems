package main

import (
	"fmt"
)

func main() {

	commands := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	num := [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}}
	runFunc(commands, num)

}

func runFunc(input []string, values [][]string) {
	var trie Trie

	for i, s := range input {
		if s == "Trie" {
			trie = Constructor()
		} else if s == "insert" {
			trie.Insert(values[i][0])
		} else if s == "search" {
			fmt.Println(trie.Search(values[i][0]))
		} else if s == "startsWith" {
			fmt.Println(trie.StartsWith(values[i][0]))
		}
	}
}

type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}
type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{root: &trieNode{}}
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
		if curr.childrens[index] == nil {
			return false // we have reached the end of a branch without getting the full word
		}
		curr = curr.childrens[index]
	}
	if curr.isWordEnd {
		return true // only return true if this endpoint represents and entire word
	}
	return false // return false if this point is not the end of a word

}

func (this *Trie) StartsWith(prefix string) bool {
	wordLength := len(prefix)
	curr := this.root

	for i := 0; i < wordLength; i++ {
		index := prefix[i] - 'a'
		if curr.childrens[index] == nil {
			return false // we have reached the end of a branch without getting the full word
		}
		curr = curr.childrens[index]
	}
	// unlike with find word we don't care if the word continues or not
	return true
}
