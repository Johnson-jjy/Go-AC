package cig

// 类似于单调栈的思想: 用另一个栈在每次调整时暂存相关元素
// 待补充:可用堆,思想类似优先队列
type SortedStack struct {
	store []int
	temp []int
}

func Constructor5() SortedStack {
	s := make([]int, 0)
	t := make([]int, 0)
	return SortedStack {
		s,
		t,
	}
}

func (this *SortedStack) Push(val int)  {
	for !this.IsEmpty() && this.store[len(this.store) - 1] < val {
		this.temp = append(this.temp, this.store[len(this.store) - 1])
		this.store = this.store[:len(this.store) - 1]
	}
	this.store = append(this.store, val)
	for len(this.temp) > 0 {
		this.store = append(this.store, this.temp[len(this.temp) - 1])
		this.temp = this.temp[:len(this.temp) - 1]
	}
}


func (this *SortedStack) Pop()  {
	if len(this.store) == 0 {
		return
	}
	this.store = this.store[:len(this.store) - 1]
}


func (this *SortedStack) Peek() int {
	if len(this.store) == 0 {
		return -1
	}
	return this.store[len(this.store) - 1]
}


func (this *SortedStack) IsEmpty() bool {
	if len(this.store) == 0 {
		return true
	}
	return false
}


/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */
