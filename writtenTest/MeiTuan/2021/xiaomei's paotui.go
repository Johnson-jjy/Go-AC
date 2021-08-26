package main

/*
5 2
5 10
8 9
1 4
7 9
6 10
*/

import (
	"fmt"
	"sort"
)

// 只过了5/10
// 原因未知
type heap []good

func (h heap)parent(i int) int {
	return i / 2
}

func (h heap)left(i int) int {
	return 2 * i
}

func (h heap)right(i int) int {
	return 2 * i + 1
}

func (h *heap)swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h heap)less(i, j int) bool {
	if h[i].value < h[j].value {
		return true
	}
	return false
}

func (h heap)len() int {
	return len(h)
}

func (h *heap)down(i int) {
	smallest := i
	left := h.left(i)
	right := h.right(i)
	if left < h.len() && h.less(left, smallest) {
		smallest = left
	}
	if right < h.len() && h.less(right, smallest) {
		smallest = right
	}
	if smallest != i {
		h.swap(i, smallest)
		h.down(smallest) // 注意这里不是i, 是对新换到的那个点做处理,同理,下面时parent
	}
}

func (h *heap)up(i int)  {
	parent := h.parent(i)
	if parent > 0 && h.less(i, parent) {
		h.swap(i, parent)
		h.up(parent)
	}
}

func (h *heap)insert(v good) {
	*h = append(*h, v)
	h.up(h.len() - 1)
}

type good struct {
	index int
	value int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if m > n {
		m = n
	}
	var w, v int
	h := make(heap, 1)
	for i := 0; i < m; i++ {
		fmt.Scan(&v, &w)
		value := v + 2 * w
		cur := good{
			i + 1,
			value,
		}
		h.insert(cur)
	}
	for i := m; i < n; i++ {
		fmt.Scan(&v, &w)
		value := v + 2 * w
		cur := good{
			i + 1,
			value,
		}
		if cur.value > h[1].value {
			h[1] = cur
			h.down(1)
		}
	}

	res := make([]int, m)
	for i := 1; i < h.len(); i++ {
		res[i - 1] = h[i].index
	}
	sort.Ints(res)
	for i := 0; i < m - 1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Print(res[m - 1])
}
