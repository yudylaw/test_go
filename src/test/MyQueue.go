package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	//用栈实现队列，只能使用栈的基本操作： push to top, peek/pop from top, size, 和 is empty
	Stack *list.List
	Head  *list.List //队列头
}

/** Initialize your data structure here. */
func Constructor4() MyQueue {
	queue := MyQueue{}
	queue.Stack = list.New()
	queue.Head = list.New()
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.Head.Len() == 0 {
		this.Head.PushFront(x)
	} else {
		this.Stack.PushFront(x)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	headItem := this.Head.Front()
	if headItem == nil {
		return -1
	}

	this.Head.Remove(headItem)
	//使用Back模拟弹出栈顶元素
	for val := this.Stack.Back(); val != nil; val = val.Prev() {
		this.Stack.Remove(val)
		this.Head.PushFront(val.Value.(int))
	}

	head := this.Head.Front()
	if head != nil {
		this.Head.Remove(head)
	} else {
		//队列为空
		return headItem.Value.(int)
	}

	for val := this.Head.Back(); val != nil; val = val.Prev() {
		this.Head.Remove(val)
		this.Stack.PushFront(val.Value.(int))
	}

	this.Head.PushFront(head.Value.(int))

	return headItem.Value.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	val := this.Head.Front()
	if val == nil {
		return -1
	}
	return val.Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Head.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor4()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	val := queue.Pop() // 返回 1
	fmt.Println(val)
	val = queue.Pop() // 返回 1
	fmt.Println(val)
	queue.Empty() // 返回 false
}
