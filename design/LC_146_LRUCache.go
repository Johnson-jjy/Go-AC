package design

import "fmt"

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */

// 解法1: 用了两个map分别存key-node和node-key,为了能在删除节点时由node删key
// 注: 将节点移动至链表头时,需要判断节点是否本就在链表中,在的话还需对节点原本的前后节点进行修改
type Node struct {
	val int
	next *Node
	pre *Node
}

type lru struct {
	size int
	cur int
	head *Node
	tail *Node
	store map[int]*Node
	forD map[*Node]int
}

func initNode(val int) *Node {
	return &Node {
		val,
		nil,
		nil,
	}
}

func initLRU(k int) *lru {
	head := initNode(-1)
	tail := initNode(-1)
	head.next = tail
	tail.pre = head
	store := make(map[int]*Node, 0)
	forD := make(map[*Node]int, 0)
	return &lru {
		k,
		0,
		head,
		tail,
		store,
		forD,
	}
}
func (l *lru)addTohead(node *Node) {
	if _, ok := l.forD[node]; ok {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	node.next = l.head.next
	l.head.next = node
	node.next.pre = node
	node.pre = l.head
}

func (l *lru)get(key int) int {
	res := -1
	if v, ok := l.store[key]; ok {
		res = v.val
		l.addTohead(v)
	}
	return res
}

func (l *lru)set(key, val int) {
	node := initNode(val)
	l.cur++
	l.addTohead(node)
	l.store[key] = node
	l.forD[node] = key
	if l.cur > l.size {
		l.delTail()
	}
}

func (l *lru)delTail() {
	if l.head.next == l.tail {
		return
	}
	fmt.Printf("%v\n", l.head.next)
	tmp := l.tail.pre
	tmp.pre.next = l.tail
	l.tail.pre = tmp.pre
	//fmt.Printf("%v\n", tmp)
	delete(l.store, l.forD[tmp]) // 这里还需要想办法保存node对应的key，以便删除
	delete(l.forD, tmp)
	l.cur--
}

func LRU( operators [][]int ,  k int ) []int {
	// write code here
	res := make([]int, 0)
	lru := initLRU(k)
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lru.set(operators[i][1], operators[i][2])
		} else if operators[i][0] == 2 {
			cur := lru.get(operators[i][1])
			res = append(res, cur)
		}
	}
	return res
}


// 解法2: 与1有亮点不同 ->
//		1.用root单个指针替换了头指针和尾指针(借鉴了go_list的实现);
//		2.在Node节点中既存key,又存val,保证了删除节点时可从Node中直接获取key以便删除map中的对应项

//type Node struct {
//	key, value int
//	prev, next *Node
//}
//
//type LRUCache struct {
//	capacity,size int
//	root *Node
//	kvMap map[int]*Node
//}
//
//
//func Constructor146(capacity int) LRUCache {
//	cur := &Node {
//		key : 0,
//		value : 0,
//	}
//	cur.next = cur
//	cur.prev = cur
//	fmt.Printf("cur: %v", cur)
//
//	return LRUCache {
//		capacity : capacity,
//		size : 0,
//		root : cur,
//		kvMap : make(map[int]*Node),
//	}
//}
//
//func (this *LRUCache) PutFront (cur *Node) {
//	if cur.prev != nil {
//		cur.prev.next = cur.next
//		cur.next.prev = cur.prev
//	}
//	this.root.next.prev = cur
//	cur.next = this.root.next
//	this.root.next = cur
//	cur.prev = this.root
//}
//
//func (this *LRUCache) Get(key int) int {
//	got, ok := this.kvMap[key]
//	if !ok {
//		return -1
//	}
//	res := got.value
//	this.PutFront(got)
//	return res
//}
//
//
//func (this *LRUCache) Put(key int, value int)  {
//	got, ok := this.kvMap[key]
//	if ok {
//		got.value = value
//		this.PutFront(got)
//	} else {
//		var cur *Node
//		cur = new(Node)
//		cur.value = value
//		cur.key = key
//		//fmt.Printf("test: %v\t%v\n", cur, this.kvMap)
//		this.PutFront(cur)
//		this.kvMap[key] = cur
//		//fmt.Printf("test: %v\t%v\n", cur, this.kvMap)
//		this.size++
//		if this.size > this.capacity {
//			node := this.DeleteTail()
//			delete(this.kvMap, node.key)
//			this.size--
//		}
//	}
//}
//
//func (this *LRUCache) DeleteTail() *Node {
//	if this.root.prev == nil {
//		return nil
//	}
//	node := this.root.prev
//	node.prev.next = this.root
//	this.root.prev = node.prev
//	return node
//}
//
//func main()  {
//	obj := Constructor146(2)
//	fmt.Printf("obj: %v\n", obj)
//	obj.Put(1, 1)
//	obj.Put(2, 2)
//	test1 := obj.Get(1)
//	fmt.Printf("%vtest1: %v\n", obj, test1)
//	obj.Put(3,3)
//	test2 := obj.Get(2)
//	fmt.Printf("%vtest2: %v\n", obj, test2)
//	obj.Put(4,4)
//	test11 := obj.Get(1)
//	fmt.Printf("%vtest11: %v\n", obj, test11)
//	test3 := obj.Get(3)
//	fmt.Printf("%vtest3: %v\n", obj, test3)
//	test4 := obj.Get(4)
//	fmt.Printf("%vtest4: %v\n", obj, test4)
//}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */