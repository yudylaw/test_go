package main

type DequeNode struct {
	Next *DequeNode
	Pre  *DequeNode
	Val  int
}

type MyCircularDeque struct {
	//设计循环双端队列
	head *DequeNode
	tail *DequeNode
	cap  int
	len  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor8(k int) MyCircularDeque {
	queue := MyCircularDeque{cap: k, len: 0}
	return queue
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.len == this.cap {
		return false
	}
	if this.head == nil {
		node := &DequeNode{Val: value}
		//node.Next = node
		//node.Pre = node
		this.head = node
		this.tail = node
		this.len = this.len + 1
		return true
	}
	tmp := this.head
	node := &DequeNode{Val: value}
	node.Next = tmp
	node.Pre = this.tail
	tmp.Pre = node
	this.tail.Next = node
	this.head = node
	this.len = this.len + 1
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.len == this.cap {
		return false
	}
	if this.tail == nil {
		node := &DequeNode{Val: value}
		//node.Next = node
		//node.Pre = node
		this.head = node
		this.tail = node
		this.len = this.len + 1
		return true
	}
	tmp := this.tail
	node := &DequeNode{Val: value}
	node.Pre = tmp
	node.Next = this.head
	tmp.Next = node
	this.head.Pre = node
	this.tail = node
	this.len = this.len + 1
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.head == nil {
		return false
	}
	if this.head == this.tail {
		this.head.Pre = nil
		this.head.Next = nil
		this.head = nil
		this.tail = nil
		this.len = this.len - 1
		return true
	}
	node := this.head.Next
	tmp := this.head
	tmp.Pre = nil
	tmp.Next = nil
	node.Pre = this.tail
	this.tail.Next = node
	this.head = node
	this.len = this.len - 1
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.tail == nil {
		return false
	}
	if this.head == this.tail {
		this.tail.Pre = nil
		this.tail.Next = nil
		this.head = nil
		this.tail = nil
		this.len = this.len - 1
		return true
	}
	node := this.tail.Pre
	tmp := this.tail
	tmp.Pre = nil
	tmp.Next = nil
	node.Next = this.head
	this.head.Pre = node
	this.tail = node
	this.len = this.len - 1
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.head == nil {
		return -1
	}
	return this.head.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.tail == nil {
		return -1
	}
	return this.tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
