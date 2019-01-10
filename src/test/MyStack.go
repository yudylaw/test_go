package main

import "container/list"

type MyStack struct {
	//用队列实现栈，只能使用队列的基本操作：push to back, peek/pop from front, size, 和 is empty
	//队列 FIFO，栈：LIFO
	Queue *list.List
}

/** Initialize your data structure here. */
func Constructor3() MyStack {
	obj := MyStack{}
	obj.Queue = list.New() //双向链表
	return obj
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Queue.PushBack(x)
	for i := 0; i < this.Queue.Len()-1; i++ {
		this.Queue.MoveToBack(this.Queue.Front())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.Queue.Front()
	if val != nil {
		this.Queue.Remove(val)
	} else {
		return -1
	}
	return val.Value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	val := this.Queue.Front()
	if val != nil {
		return val.Value.(int)
	}
	return -1
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
