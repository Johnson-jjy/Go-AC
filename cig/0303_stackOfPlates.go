package cig

type StackOfPlates struct {
	store [][]int
	single int
}

func Constructor3(cap int) StackOfPlates {
	s := make([][]int, 0)
	return StackOfPlates {
		s,
		cap,
	}
}

func (this *StackOfPlates) Push(val int)  {
	if this.single == 0 {
		return
	}
	n := len(this.store) - 1
	if  n == -1 || len(this.store[n]) == this.single {
		cur := make([]int, 1)
		cur[0] = val
		this.store = append(this.store, cur)
	} else {
		this.store[n] = append(this.store[n], val)
	}
}


func (this *StackOfPlates) Pop() int {
	n := len(this.store) - 1
	if n == -1 {
		return -1
	}
	cur := this.store[n]
	m := len(cur) - 1
	res := cur[m]
	if m == 0 {
		this.store = this.store[:n]
	} else {
		this.store[n] = (this.store[n])[:m]
	}
	return res
}


func (this *StackOfPlates) PopAt(index int) int {
	n := len(this.store) - 1
	if n == -1 || index > n {
		return -1
	}
	if n == index {
		return this.Pop()
	}
	cur := this.store[index]
	m := len(cur) - 1
	res := cur[m]
	if m == 0 {
		this.store = append(this.store[:index], append([][]int{}, this.store[index+1:]...)...)
	} else {
		this.store[index] = (this.store[index])[:m]
	}
	return res
}

//if len(plate) == 0 {
//tmp := this.Stack[index+1:]
//this.Stack = this.Stack[:index]
//this.Stack = append(this.Stack, tmp...)
//}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
