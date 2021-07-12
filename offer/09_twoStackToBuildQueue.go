package offer

type CQueue struct {
	stack1 []int
	stack2 []int
}


func Constructor() CQueue {
	stack1 := make([]int, 0)
	stack2 := make([]int, 0)
	return CQueue{
		stack1,
		stack2,
	}
}


func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		if len(this.stack1) != 0 {
			for len(this.stack1) != 0 {
				this.stack2 = append(this.stack2, this.stack1[len(this.stack1) - 1])
				this.stack1 = this.stack1[:len(this.stack1) - 1]
			}
		}
	}
	if len(this.stack2) != 0 {
		res := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return res
	}
	return -1
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
