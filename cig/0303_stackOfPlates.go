package cig

// 理解题意:对于PopAt操作,需要用到切片拼接,注意相关写法
// cap为0时,需要另做处理
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
	if this.single == 0 { // cap为0, 直接返回, 不做其他处理
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
		// 对于临时append的使用,第一个参数依然需要是高维,两次...降维
		this.store = append(this.store[:index], append([][]int{}, this.store[index+1:]...)...) // 注意拼接的写法
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
