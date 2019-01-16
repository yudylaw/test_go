package main

import "fmt"

type TrieNode struct {
	//前缀树（字典树）
	Nodes  []*TrieNode
	Word   string
	IsNode bool
}
type Trie struct {
	Root *TrieNode
}

/** Initialize your data structure here. */
func Constructor9() Trie {
	nodes := make([]*TrieNode, 0)
	node := &TrieNode{Nodes: nodes}
	trie := Trie{Root: node}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var node *TrieNode = this.Root
	for i := 1; i <= len(word) && node != nil; i++ {
		w := word[:i]
		flag := false
		for _, n := range node.Nodes {
			if n.Word == w {
				flag = true
				node = n
				break
			}
		}
		isNode := i == len(word)
		if flag {
			if !node.IsNode && isNode {
				node.IsNode = true
			}
		} else {
			n := &TrieNode{Word: w, IsNode: isNode}
			if node.Nodes == nil {
				node.Nodes = make([]*TrieNode, 0)
			}
			node.Nodes = append(node.Nodes, n)
			node = n
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this.Root == nil || len(word) == 0 {
		return false
	}
	var node *TrieNode = this.Root
	for i := 1; i <= len(word) && node != nil; i++ {
		w := word[:i]
		flag := false
		for _, n := range node.Nodes {
			if n.Word == w {
				flag = true
				node = n
				break
			}
		}
		if flag && i == len(word) && node.IsNode {
			return true
		}
		if !flag {
			return false
		}
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this.Root == nil {
		return false
	}
	var node *TrieNode = this.Root
	for i := 1; i <= len(prefix) && node != nil; i++ {
		w := prefix[:i]
		flag := false
		for _, n := range node.Nodes {
			if n.Word == w {
				flag = true
				node = n
				break
			}
		}
		//prefix
		if flag && i == len(prefix) {
			return true
		}
		if !flag {
			return false
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor9()
	trie.Insert("apple")
	flag := trie.Search("apple") // 返回 true
	fmt.Println(flag)
	flag = trie.Search("app") // 返回 false
	fmt.Println(flag)
	flag = trie.StartsWith("app") // 返回 true
	fmt.Println(flag)
	trie.Insert("app")
	flag = trie.Search("app") // 返回 true
	fmt.Println(flag)
}
