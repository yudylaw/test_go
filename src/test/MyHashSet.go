package main

import (
	"container/list"
)

type MyHashSet struct {
	keys       []int
	len        int
	cap        int
	loadFactor float32
	linkedList []*list.List
}

const NIL_KEY int = -1

/** Initialize your data structure here. */
func Constructor6() MyHashSet {
	hashSet := MyHashSet{len: 0, loadFactor: 0.75}
	capacity := 16
	hashSet.keys = make([]int, capacity)
	for i, _ := range hashSet.keys {
		hashSet.keys[i] = NIL_KEY
	}
	hashSet.linkedList = make([]*list.List, capacity)
	hashSet.cap = capacity
	return hashSet
}

func (this *MyHashSet) hashCode(key int) int {
	//模仿Java：高16bit不变，低16bit和高16bit做了一个异或
	h := key ^ (key >> 16)
	return h
}

func (this *MyHashSet) resize() {
	//扩容2倍
	newKeys := make([]int, 2*this.cap)
	for i, _ := range newKeys {
		newKeys[i] = NIL_KEY
	}
	newList := make([]*list.List, 2*this.cap)
	this.cap = 2 * this.cap
	//rehash
	for i, key := range this.keys {
		if key == NIL_KEY {
			continue
		}
		j := this.getIndex(key)
		if newKeys[j] == NIL_KEY {
			newKeys[j] = key
		} else {
			if newKeys[j] != key {
				//冲突
				if newList[j] == nil {
					newList[j] = list.New()
				}
				newList[j].PushBack(key)
			}
		}
		nodes := this.linkedList[i]
		if nodes != nil {
			for node := nodes.Front(); node != nil; node = node.Next() {
				key := node.Value.(int)
				k := this.getIndex(key)
				if newKeys[k] == NIL_KEY {
					//插入新位置
					newKeys[k] = key
				} else {
					if newKeys[k] != key {
						//冲突
						if newList[k] == nil {
							newList[k] = list.New()
						}
						newList[k].PushBack(key)
					}
				}
			}
			this.linkedList[i] = nil
		}
	}
	this.keys = newKeys
	this.linkedList = newList
	//释放存储 TODO
}

func (this *MyHashSet) getIndex(key int) int {
	hash := this.hashCode(key)
	return hash % this.cap
}

func (this *MyHashSet) Add(key int) {
	if this.len >= int(this.loadFactor*float32(this.cap)) {
		this.resize()
	}
	index := this.getIndex(key)
	k := this.keys[index]
	if k == NIL_KEY {
		//直接插入
		this.keys[index] = key
		this.len = this.len + 1
		return
	}
	if k == key {
		//已存在
	} else {
		//key冲突
		nodes := this.linkedList[index]
		if nodes == nil {
			this.linkedList[index] = list.New()
			//插入队尾
			this.linkedList[index].PushBack(key)
			this.len = this.len + 1
		} else {
			for node := nodes.Front(); node != nil; node = node.Next() {
				if node.Value.(int) == key {
					//已存在
					return
				}
			}
			//插入队尾
			nodes.PushBack(key)
			this.len = this.len + 1
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.getIndex(key)
	k := this.keys[index]
	if k == NIL_KEY {
		return false
	}
	if k == key {
		return true
	} else {
		//key冲突
		nodes := this.linkedList[index]
		if nodes != nil {
			for node := nodes.Front(); node != nil; node = node.Next() {
				if node.Value.(int) == key {
					return true
				}
			}

		}
		return false
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.getIndex(key)
	k := this.keys[index]
	if k == NIL_KEY {
		return
	}
	if k == key {
		nodes := this.linkedList[index]
		if nodes != nil {
			//存在链表
			node := nodes.Front()
			//移动节点
			this.keys[index] = node.Value.(int)
			nodes.Remove(node)
			if nodes.Len() == 0 {
				this.linkedList[index] = nil
			}
		} else {
			//不存在链表
			this.keys[index] = NIL_KEY
		}
		this.len = this.len - 1
	} else {
		//key冲突
		nodes := this.linkedList[index]
		if nodes != nil {
			for node := nodes.Front(); node != nil; node = node.Next() {
				if node.Value.(int) == key {
					nodes.Remove(node)
					if nodes.Len() == 0 {
						this.linkedList[index] = nil
					}
					this.len = this.len - 1
					return
				}
			}

		}
	}
}
