package main

import (
	"container/list"
	"fmt"
)

type MyHashMap struct {
	keys       []int
	values     []int
	len        int
	cap        int
	linkedList []*list.List
}

type KV struct {
	Key   int
	Value int
}

/** Initialize your data structure here. */
func Constructor5() MyHashMap {
	hashMap := MyHashMap{len: 0}
	capacity := 16
	hashMap.keys = make([]int, capacity)
	hashMap.values = make([]int, capacity)
	//for i, _ := range hashMap.keys {
	//	hashMap.keys[i] = -1
	//}
	hashMap.cap = capacity
	return hashMap
}

func (this *MyHashMap) hashCode(key int) int {
	//模仿Java：高16bit不变，低16bit和高16bit做了一个异或
	h := key ^ (key >> 16)
	return h
}

func (this *MyHashMap) resize() {
	//扩容2倍
	newKeys := make([]int, 2*this.cap)
	newValues := make([]int, 2*this.cap)
	//for i, _ := range newKeys {
	//	newKeys[i] = -1
	//}
	copy(newKeys, this.keys)
	copy(newValues, this.values)
	this.cap = 2 * this.cap
	this.keys = newKeys
	this.values = newValues
}

func (this *MyHashMap) getIndex(key int) int {
	hash := this.hashCode(key)
	return hash % this.cap
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := this.getIndex(key)
	k := this.keys[index]
	if k == key {
		if this.len == this.cap {
			this.resize()
		}
		//rehash
		index = this.getIndex(key)
		this.values[index] = value
		return
	} else {
		//key冲突
		nodes := this.linkedList[index]
		if nodes == nil {
			this.linkedList[index] = list.New()
			//插入队尾
			kv := KV{Key: key, Value: value}
			this.linkedList[index].PushBack(kv)
		} else {
			for node := nodes.Front(); node != nil; node = node.Next() {
				kv := node.Value.(KV)
				if kv.Key == key {
					//更新
					kv.Value = value
					node.Value = kv
					return
				}
			}
			//插入队尾
			kv := KV{Key: key, Value: value}
			nodes.PushBack(kv)
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := this.getIndex(key)
	k := this.keys[index]
	if k == 0 {
		return -1
	}
	if k == key {
		return this.values[index]
	} else {
		//key冲突
		nodes := this.linkedList[index]
		if nodes != nil {
			for node := nodes.Front(); node != nil; node = node.Next() {
				kv := node.Value.(KV)
				if kv.Key == key {
					return kv.Value
				}
			}

		}
		return -1
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := this.getIndex(key)
	k := this.keys[index]
	if k == 0 {
		return
	}
	if k == key {
		//BUG 存在链表的情况如何处理
		this.keys[index] = 0
		this.values[index] = 0
	} else {
		//key冲突

	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	obj := Constructor5()
	obj.Put(1, 2)
	value := obj.Get(1)
	fmt.Printf("value=%d\n", value)
	obj.Remove(1)
	value = obj.Get(2)
	fmt.Printf("value=%d\n", value)

}
