package main

import "fmt"

//堆排序 -- 本质上可以理解为建堆之后的出堆过程
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}

// 类似于down的过程, 堆顶为最大元素
func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		// 每次对当前长度的arr进行堆化操作(最大元素在堆顶)
		HeapSortMax(arr, lastlen)
		if i < length { // 交换堆顶(max)元素到队列尾部
			// 交换之后的arr非最大堆
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}
