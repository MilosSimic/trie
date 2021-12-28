package main

import (
	"fmt"
)

type Trie struct {
	isEnd bool
	children map[rune]*Trie
}

func NewTrie() *Trie {
	return &Trie{
		children:map[rune]*Trie{},
		isEnd: false,
	}
}

func (t *Trie) Insert(data string) {
	parent := t
	for _, ch := range data {
		if child, ok := parent.children[ch]; ok { // if char exists, than look up in level under it
			parent = child
		}else{
			newChild := &Trie{
				isEnd:false,
				children:map[rune]*Trie{},
			}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isEnd = true
}

func (t *Trie) Lookup(data string) bool {
	parent := t
	for _, ch := range data {
		if child, ok := parent.children[ch]; ok { // if char exists, than look up in level under it
			parent = child
		}else{
			return false
		}
	}
	return parent.isEnd
}

func (t *Trie) SartsWith(prefix string) bool {
	parent := t
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok { // if char exists, than look up in level under it
			parent = child
		}else{
			return false
		}
	}
	return true
}

func (t *Trie) SartsWithCount(prefix string) int {
	parent := t
	counter := 0
	for _, ch := range prefix {
		if child, ok := parent.children[ch]; ok { // if char exists, than look up in level under it
			parent = child
			counter++
		}else{
			return 0
		}
	}
	return counter
}

func main(){
	trie := NewTrie()
	trie.Insert("Hello")
	trie.Insert("World")
	trie.Insert("Girl")
	trie.Insert("Gigi")
	fmt.Println(trie.Lookup("Hello"))
	fmt.Println(trie.Lookup("Man"))
	fmt.Println(trie.SartsWith("Gi"))
	fmt.Println(trie.SartsWith("Ot"))
	fmt.Println(trie.SartsWithCount("Gi"))
}
