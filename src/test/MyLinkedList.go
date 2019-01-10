package main

import "fmt"

type ListedNode struct {
	Val  int
	Next *ListedNode
}

type MyLinkedList struct {
	Head *ListedNode
	Tail *ListedNode
	Len  int
}

/** Initialize your data structure here. */
func Constructor2() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	for i, cur := 0, this.Head; cur != nil && i <= index; cur, i = cur.Next, i+1 {
		if i == index {
			return cur.Val
		}
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	tmp := this.Head
	node := &ListedNode{Val: val, Next: tmp}
	this.Head = node
	this.Len = this.Len + 1
	if this.Tail == nil {
		this.Tail = this.Head
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListedNode{Val: val}
	tmp := this.Tail
	if tmp != nil {
		tmp.Next = node
	}
	this.Tail = node
	this.Len = this.Len + 1
	if this.Head == nil {
		this.Head = this.Tail
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Len {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Len {
		this.AddAtTail(val)
		return
	}
	pre := this.Head
	for i, cur := 0, this.Head; cur != nil; cur, i = cur.Next, i+1 {
		if i == index {
			node := &ListedNode{Val: val, Next: cur}
			pre.Next = node
			this.Len = this.Len + 1
			return
		}
		pre = cur
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Len {
		return
	}
	pre := this.Head
	for i, cur := 0, this.Head; cur != nil; cur, i = cur.Next, i+1 {
		if i == index {
			if pre != nil {
				pre.Next = cur.Next
			}
			cur.Next = nil
			this.Len = this.Len - 1
			return
		}
		pre = cur
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	linkedList := MyLinkedList{}
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	val := linkedList.Get(1)
	fmt.Printf("value of index 1=%d\n", val)
	linkedList.DeleteAtIndex(1)
	val = linkedList.Get(1)
	fmt.Printf("value of index 1=%d", val)
}
