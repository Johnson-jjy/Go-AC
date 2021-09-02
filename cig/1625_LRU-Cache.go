package cig

// 经典LRU-Cache
// 关于删除tail的两种思路: 1.Node中保存key值; 2.LRU的Cache中保存一个nodeKey的map
// 注意: 删除节点或添加节点时,一定要把该添加到位的全都添加到位
type LRUCache struct {
	size int
	curLen int
	head *Node
	tail *Node
	keyNode map[int]*Node
}

type Node struct {
	key int
	val int
	next *Node
	pre *Node
}

func Init(key int, val int) *Node {
	return &Node{
		key,
		val,
		nil,
		nil,
	}
}

func Constructor1625(capacity int) LRUCache {
	head := Init(-1, -1)
	tail := Init(-1,- 1)
	head.next = tail
	tail.pre = head
	kN := make(map[int]*Node, 0)
	return LRUCache{
		capacity,
		0,
		head,
		tail,
		kN,
	}
}

func (this *LRUCache) MoveToHead(key int) {
	cur := this.keyNode[key]
	cur.pre.next = cur.next
	cur.next.pre = cur.pre
	cur.next = this.head.next
	cur.next.pre = cur
	this.head.next = cur
	cur.pre = this.head
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.keyNode[key]; ok {
		this.MoveToHead(key)
		return v.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.keyNode[key]; ok {
		this.MoveToHead(key)
		this.keyNode[key].val = value
	} else {
		this.AddOne(key, value)
	}
}

func (this *LRUCache) AddOne(key, val int) {
	node := Init(key, val)
	this.head.next.pre = node
	node.next = this.head.next
	this.head.next = node
	node.pre = this.head
	this.curLen++
	this.keyNode[key] = node
	if this.curLen > this.size {
		this.DeleteTail()
	}
}

func (this *LRUCache) DeleteTail() {
	tail := this.tail.pre
	tail.pre.next = tail.next
	tail.next.pre = tail.pre
	tailKey := tail.key
	delete(this.keyNode, tailKey)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
