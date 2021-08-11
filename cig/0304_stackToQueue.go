package cig

// 两栈互相倒
type MyQueue struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor4() MyQueue {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	return MyQueue {
		s1,
		s2,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.stack2) != 0 {
		res := this.stack2[len(this.stack2) - 1]
		this.stack2 = this.stack2[:len(this.stack2) - 1]
		return res
	}
	for len(this.stack1) != 0 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
	if len(this.stack2) == 0 {
		return -1
	}
	res := this.stack2[len(this.stack2) - 1]
	this.stack2 = this.stack2[:len(this.stack2) - 1]
	return res
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.stack2) != 0 {
		return this.stack2[len(this.stack2) - 1]
	}
	for len(this.stack1) != 0 {
		this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
		this.stack1 = this.stack1[:len(this.stack1) - 1]
	}
	if len(this.stack2) == 0 {
		return -1
	}
	return this.stack2[len(this.stack2) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stack1) == 0 && len(this.stack2) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
