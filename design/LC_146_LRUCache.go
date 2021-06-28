package design

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity,size int
	root *Node
	kvMap map[int]*Node
}


func Constructor146(capacity int) LRUCache {
	cur := &Node {
		key : 0,
		value : 0,
	}
	cur.next = cur
	cur.prev = cur
	fmt.Printf("cur: %v", cur)

	return LRUCache {
		capacity : capacity,
		size : 0,
		root : cur,
		kvMap : make(map[int]*Node),
	}
}

func (this *LRUCache) PutFront (cur *Node) {
	if cur.prev != nil {
		cur.prev.next = cur.next
		cur.next.prev = cur.prev
	}
	this.root.next.prev = cur
	cur.next = this.root.next
	this.root.next = cur
	cur.prev = this.root
}

func (this *LRUCache) Get(key int) int {
	got, ok := this.kvMap[key]
	if !ok {
		return -1
	}
	res := got.value
	this.PutFront(got)
	return res
}


func (this *LRUCache) Put(key int, value int)  {
	got, ok := this.kvMap[key]
	if ok {
		got.value = value
		this.PutFront(got)
	} else {
		var cur *Node
		cur = new(Node)
		cur.value = value
		cur.key = key
		//fmt.Printf("test: %v\t%v\n", cur, this.kvMap)
		this.PutFront(cur)
		this.kvMap[key] = cur
		//fmt.Printf("test: %v\t%v\n", cur, this.kvMap)
		this.size++
		if this.size > this.capacity {
			node := this.DeleteTail()
			delete(this.kvMap, node.key)
			this.size--
		}
	}
}

func (this *LRUCache) DeleteTail() *Node {
	if this.root.prev == nil {
		return nil
	}
	node := this.root.prev
	node.prev.next = this.root
	this.root.prev = node.prev
	return node
}

func main()  {
	obj := Constructor146(2)
	fmt.Printf("obj: %v\n", obj)
	obj.Put(1, 1)
	obj.Put(2, 2)
	test1 := obj.Get(1)
	fmt.Printf("%vtest1: %v\n", obj, test1)
	obj.Put(3,3)
	test2 := obj.Get(2)
	fmt.Printf("%vtest2: %v\n", obj, test2)
	obj.Put(4,4)
	test11 := obj.Get(1)
	fmt.Printf("%vtest11: %v\n", obj, test11)
	test3 := obj.Get(3)
	fmt.Printf("%vtest3: %v\n", obj, test3)
	test4 := obj.Get(4)
	fmt.Printf("%vtest4: %v\n", obj, test4)
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
