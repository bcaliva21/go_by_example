package main

import "fmt"

const ALPHA_SIZE = 26

type Trie struct {
	Children [ALPHA_SIZE]*Trie
	IsWord bool
}

func (t *Trie) Insert (word string) {
	curr := t
	for _, ch := range word {
		idx := ch - 'a'
		if curr.Children[idx] == nil {
			curr.Children[idx] = &Trie{}
		}
		curr = curr.Children[idx]
	}
	curr.IsWord = true
}

func (t *Trie) Search (word string) bool {
	curr := t
	for _, ch := range word {
		idx := ch - 'a'
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return curr.IsWord
}

func (t *Trie) StartsWith (prefix string) bool {
	curr := t
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.Children[idx] == nil {
			return false
		}
		curr = curr.Children[idx]
	}
	return true
}

func main() {
	myTrie := &Trie{}

	myTrie.Insert("help")
	myTrie.Insert("held")
	myTrie.Insert("helmet")
	myTrie.Insert("helicopter")
	
	fmt.Println(myTrie)

	fmt.Println(myTrie.Search("hello"))
	fmt.Println(myTrie.Search("help"))
	fmt.Println(myTrie.StartsWith("he"))
	fmt.Println(myTrie.StartsWith("hel"))
}

