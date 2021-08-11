package cig

// 两个栈存: 正常栈和min栈,min栈可每次都一起取放值
type MinStack struct {
	store []int
	min []int
	size int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	store := make([]int, 0)
	min := make([]int, 0)
	size := 0
	return MinStack {
		store,
		min,
		size,
	}
}


func (this *MinStack) Push(x int)  {
	this.store = append(this.store, x)
	if len(this.min) == 0 || this.min[len(this.min) - 1] > x {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[len(this.min) - 1])
	}
}


func (this *MinStack) Pop()  {
	this.store = this.store[:len(this.store) - 1]
	this.min = this.min[:len(this.min) - 1]
}


func (this *MinStack) Top() int {
	if len(this.store) == 0 {
		return -1
	}
	return this.store[len(this.store) - 1]
}


func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return - 1
	}
	return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
