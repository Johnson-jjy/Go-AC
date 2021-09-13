package main

import (
	"container/heap"
	"fmt"
)

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h intHeap) Swap(i, j int)  {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{})  {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	res := old[n - 1]
	*h = old[0 : n - 1]
	return res
}

func main()  {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var k int
	fmt.Scan(&k)
	if n < k {
		k = n
	}
	h := &intHeap{}
	heap.Init(h)

	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}
	for i := k; i < n; i++ {
		heap.Push(h, arr[i])
		heap.Pop(h)
	}
	store := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		store[i] = heap.Pop(h).(int)
	}

	for i := 0; i < k; i++ {
		fmt.Println(store[i])
	}
}
