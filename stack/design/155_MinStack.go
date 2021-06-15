package design

type MinStack struct {
	normalStack []int
	minStack []int
}


/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack {
		normalStack : make([]int, 0),
		minStack : make([]int, 0),
	}
}


func (this *MinStack) Push(val int)  {
	this.normalStack = append(this.normalStack, val)
	if len(this.minStack) > 0 {
		cur := this.minStack[len(this.minStack) - 1]
		if cur < val {
			this.minStack = append(this.minStack, cur)
			return
		}
	}
	this.minStack = append(this.minStack, val)
}


func (this *MinStack) Pop()  {
	this.normalStack = this.normalStack[ : len(this.normalStack) - 1]
	this.minStack = this.minStack[ : len(this.minStack) - 1]
}


func (this *MinStack) Top() int {
	return this.normalStack[len(this.normalStack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */