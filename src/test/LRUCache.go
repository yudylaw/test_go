package main

type LinkNode struct {
	Pre  *LinkNode
	Next *LinkNode
	Val  int
}

type LRUCache struct {
	values   map[int]int
	head     *LinkNode
	tail     *LinkNode
	capacity int
	len      int
}

func Constructor(capacity int) LRUCache {
	if capacity <= 0 {
		panic("capacity must > 0")
	}
	cache := LRUCache{capacity: capacity}
	cache.values = make(map[int]int)
	return cache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.values[key]; ok {
		this.moveToHead(key)
		return val
	} else {
		//fmt.Println("key not found")
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.values[key]; ok {
		this.moveToHead(key)
	} else {
		//已满
		if this.len == this.capacity {
			this.evictTail()
		} else {
			this.len = this.len + 1
		}
		this.putToHead(key)
	}
	this.values[key] = value
}

func (this *LRUCache) evictTail() {
	//fmt.Println("evictTail")
	if this.tail == nil {
		return
	}
	lastNode := this.tail.Pre
	tailNode := this.tail
	key := tailNode.Val
	if lastNode != nil {
		lastNode.Next = nil
		tailNode.Pre = nil
		this.tail = lastNode
	} else {
		this.tail = nil
		this.head = nil
	}
	//fmt.Printf("evict key=%d\n", key)
	delete(this.values, key)
}

func (this *LRUCache) putToHead(key int) {
	//fmt.Println("putToHead, key=%d", key)
	if this.head == nil {
		this.head = &LinkNode{Val: key}
		this.tail = this.head
		return
	}
	tmp := this.head
	node := &LinkNode{Val: key, Next: tmp}
	tmp.Pre = node
	this.head = node
}

func (this *LRUCache) moveToHead(key int) {
	//fmt.Println("moveToHead, key=%d", key)
	for cur := this.head; cur != nil; cur = cur.Next {
		if cur.Val == key {
			if cur == this.head {
				return
			}
			cur.Pre.Next = cur.Next
			if cur.Next != nil {
				cur.Next.Pre = cur.Pre
			} else {
				//cur is tail node
				lastNode := this.tail.Pre
				tailNode := this.tail
				lastNode.Next = nil
				tailNode.Pre = nil
				this.tail = lastNode
			}

			tmp := this.head
			cur.Next = tmp
			tmp.Pre = cur

			this.head = cur
			return
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(1)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回  1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	//cache.Get(2)    // 返回 -1 (未找到)
	//cache.Put(4, 4) // 该操作会使得密钥 1 作废
	//cache.Get(1)    // 返回 -1 (未找到)
	//cache.Get(3)    // 返回  3
	//cache.Get(4)    // 返回  4
}
