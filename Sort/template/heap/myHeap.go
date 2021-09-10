package main

import "fmt"

// 封装一个堆
type Heap []int

func (h Heap)parent(i int) int {
	return i / 2
}

func (h Heap)left(i int) int {
	return 2 * i
}

func (h Heap)right(i int) int {
	return 2 * i + 1
}

func (h Heap)swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap)len() int {
	return len(h)
}

// 比大小 -> 若为小根堆则改为less即可
func (h Heap)greater(i, j int) bool {
	if h[i] > h[j] {
		return true
	}
	return false
}

func (h *Heap)down(i int) {
	largest := i
	left := h.left(i)
	right := h.right(i)
	if left < h.len() && h.greater(left, largest) {
		largest = left
	}
	// 注意! 这里是if,要走这个判断,而不是else-if
	if right < h.len() && h.greater(right, largest) {
		largest = right
	}

	if largest != i {
		h.swap(i, largest)
		h.down(largest)
	}
}

func (h *Heap)up(i int) {
	parent := h.parent(i)
	if parent > 0 && h.greater(i, parent) {
		h.swap(parent, i)
		h.up(parent)
	}
}

// 向堆中添加值
func (h *Heap)Insert(key int) {
	*h = append(*h, key)
	h.up(h.len() - 1)
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap)Remove(i int) (int, bool) {
	if i <= 0 || i > h.len() {
		return -1, false
	}
	ans := (*h)[i]
	h.swap(i, h.len() - 1)
	*h = (*h)[:h.len() - 1]
	h.down(i)
	// *h = (*h)[:h.len() - 1] 不能写在这里,这里起不到删除效果
	return ans, true
}

// 弹出堆顶元素元素并返回其值
func (h *Heap)Pop() int {
	ans := (*h)[1]
	h.swap(1, h.len() - 1)
	*h = (*h)[:h.len() - 1] // 同Remove理,此处顺序不能放下面
	h.down(1)
	return ans
}

// 根据任意数组初始化
func Init(arr []int) Heap{
	h := make(Heap, len(arr) + 1)
	for i := 1; i <= len(arr); i++ {
		h[i] = arr[i - 1]
	}
	//fmt.Println(h)
	for i := h.len()/2; i >= 1; i-- {
		h.down(i)
		//fmt.Println(i, h)
	}
	return h
}

func main() {
	arr := []int{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h := Init(arr)
	fmt.Println(h) // [0 30 19 25 17 15 20 3 7 10]

	h.Insert(26)
	fmt.Println(h) // [0 30 26 25 17 19 20 3 7 10 15]

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h) // 19 true [0 30 26 25 17 15 20 3 7 10]

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h) // 30 true [0 26 17 25 10 15 20 3 7]

	z := h.Pop()
	fmt.Println(z, h) // 26 [0 25 17 20 10 15 7 3]
}