package main

import "fmt"

type SubfixTreeNode struct {
	//后缀树
	Nodes  []*SubfixTreeNode
	Index  int
	Word   string
	IsWord bool
}

type SubfixTree struct {
	Root *SubfixTreeNode
	Word string
}

func Constructor10(word string) SubfixTree {
	rootNode := &SubfixTreeNode{}
	tree := SubfixTree{Root: rootNode, Word: word}
	tree.Root.Nodes = make([]*SubfixTreeNode, 0)

	words := make(map[string]*SubfixTreeNode)
	for i := 1; i <= len(word); i++ {
		w := word[i-1 : i]
		if node, ok := words[w]; ok {
			indexNode := &SubfixTreeNode{Index: i - 1}
			node.Nodes = append(node.Nodes, indexNode)
		} else {
			wordNode := &SubfixTreeNode{Word: w, IsWord: true}
			words[w] = wordNode
			tree.Root.Nodes = append(tree.Root.Nodes, wordNode)

			wordNode.Nodes = make([]*SubfixTreeNode, 0)
			indexNode := &SubfixTreeNode{Index: i - 1}
			wordNode.Nodes = append(wordNode.Nodes, indexNode)
		}
	}

	return tree
}

func (this *SubfixTree) Contains(word string) bool {
	if len(word) == 0 || word == this.Word {
		return true
	}
	if len(word) > len(this.Word) {
		return false
	}
	w := word[0:1]
	//所有后缀字符串
	indexs := make([]int, 0)
	flag := false
	for _, n := range this.Root.Nodes {
		if n.Word == w {
			for _, node := range n.Nodes {
				indexs = append(indexs, node.Index)
			}
			flag = true
			break
		}
	}
	if !flag {
		return false
	}
	for _, startIndex := range indexs {
		subfix := this.Word[startIndex:]
		if len(word) > len(subfix) {
			continue
		}
		for i := 0; i < len(word); i++ {
			//后缀字符串的前缀，满足
			if subfix[i] == word[i] && i == len(word)-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	tree := Constructor10("abacd")
	flag := tree.Contains("b")
	fmt.Println(flag)
}
