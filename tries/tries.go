package main

import "fmt"

const alphabetSize = 26

// trie is an ordered data structure where we will store the words,
// by one by one letter based on indexing. each node will contain 26 pointers
// which will store address of word's each letter

type node struct {
	children [alphabetSize]*node
	data     string
	isEnd    bool
}

type trie struct {
	root *node
}

func (t *trie) insert(word string) {
	child := t.root
	l := len(word)
	for i := 0; i < l; i++ {
		childIndex := word[i] - 'a'
		if child.children[childIndex] == nil {
			child.children[childIndex] = &node{data: string(word[i])}
		}
		child = child.children[childIndex]
	}
	child.isEnd = true

}

func (t *trie) search(word string) bool {
	child := t.root
	l := len(word)
	for i := 0; i < l; i++ {
		childIndex := word[i] - 'a'
		if child.children[childIndex] == nil {
			return false
		}
		child = child.children[childIndex]
	}
	if child.isEnd == true {
		return true
	}
	return false

}

func main() {
	t := &trie{root: &node{}}
	t.insert("aragon")
	t.insert("eragon")
	t.insert("dragon")
	fmt.Println(t.search("eragon"))
}
