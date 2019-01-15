package main

type Node struct {
	Val  int
	Next *Node
}

type MyCircularQueue struct {
	head *Node
	tail *Node
	cap  int
	len  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor7(k int) MyCircularQueue {
	queue := MyCircularQueue{cap: k, len: 0}
	return queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.len == this.cap {
		return false
	}
	if this.tail == nil {
		node := &Node{Val: value}
		this.tail = node
		this.head = node
		this.len = this.len + 1
		return true
	}
	tmp := this.tail
	node := &Node{Val: value, Next: this.head}
	tmp.Next = node
	this.tail = node
	this.len = this.len + 1
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == nil {
		return false
	}
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
		this.len = this.len - 1
		return true
	}
	tmp := this.head
	node := tmp.Next
	tmp.Next = nil
	this.head = node
	this.len = this.len - 1
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head == nil {
		return -1
	}
	return this.head.Val
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail == nil {
		return -1
	}
	return this.tail.Val
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.len == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.len == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
