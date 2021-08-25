package monotonic_stack

// 股票价格跨度
// 设计类题目,但本质思想依然是单调栈
// 难点: 本题没有数组,每次调用next只放入一个元素还要求出结果->无数组的最大难点便在于无坐标,于是我们可以自己利用sum保存坐标
// 注 -- 1.每次刚进入时对sum自增,便于计数; 2.注意题目需要的逻辑--之前小于等于,故当前元素进栈时,需大于等于
type StockSpanner struct {
	index []int
	value []int
	sum int
}


func Constructor() StockSpanner {
	sI := make([]int, 0)
	sV := make([]int, 0)
	return StockSpanner {
		sI,
		sV,
		0,
	}
}


func (this *StockSpanner) Next(price int) int {
	this.sum++
	for len(this.index) > 0 && price >= this.value[len(this.value) - 1]{
		this.index = this.index[:len(this.index) - 1]
		this.value = this.value[:len(this.value) - 1]
	}
	this.value = append(this.value, price)
	this.index = append(this.index, this.sum)
	if len(this.index) == 1 {
		return this.sum
	}
	return this.sum - this.index[len(this.index) - 2]
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
