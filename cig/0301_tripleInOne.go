package cig

// 本质上的思想: %区分不同的栈,并保存各栈所用的量即可
// 注:其实可以完全用一个栈,只是不同的位置赋予不同的意义即可: 存size的,存各栈index的,存数据的
type TripleInOne struct {
	store []int
	size int
}

func Constructor(stackSize int) TripleInOne {
	stack := make([]int, stackSize*3 + 3)
	copy(stack[:3], []int{-1, -1, -1})
	return TripleInOne {
		stack,
		stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int)  {
	if this.store[stackNum] == this.size - 1{
		return
	}
	this.store[stackNum]++
	index := (this.store[stackNum] + 1) * 3 + stackNum
	this.store[index] = value
}


func (this *TripleInOne) Pop(stackNum int) int {
	if this.store[stackNum] == -1 {
		return -1
	}
	res := this.store[(this.store[stackNum] + 1) * 3 + stackNum]
	this.store[stackNum]--
	return res
}


func (this *TripleInOne) Peek(stackNum int) int {
	if this.store[stackNum] == -1 {
		return -1
	}
	res := this.store[(this.store[stackNum] + 1) * 3 + stackNum]
	return res
}


func (this *TripleInOne) IsEmpty(stackNum int) bool {
	if this.store[stackNum] == -1 {
		return true
	}
	return false
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
