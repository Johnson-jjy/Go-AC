package main

import "fmt"

var count int

func main() {
	num := make([]int, 21)
	num[0] = 1
	for i := 1; i <= 20; i++ {
		num[i] = num[i - 1] * 2
	}

	var length int
	fmt.Scan(&length)
	originArray := make([]int, num[length])
	for i := 0; i < num[length]; i++ {
		fmt.Scan(&originArray[i])
	}

	var m int
	fmt.Scan(&m)
	time := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&time[i])
	}

	store := make(map[string]int, 0)
	for i := 0; i < m; i++ {
		every := num[time[i]]
		if every > 1 {
			originArray = reversePair(originArray, every)
		}
		cur := make([]byte, len(originArray))
		for j := 0; j < len(cur); j++ {
			cur[j] = byte(originArray[j] + '0')
		}
		str := string(cur)
		if _, ok := store[str]; !ok {
			curCount := countPairs(originArray)
			store[str] = curCount
		}
		fmt.Println(store[str])
	}
}

func reversePair(arr []int, n int) []int {
	for i := 0; i < len(arr); i = i + n {
		j := i
		k := i + n - 1
		for j < k {
			arr[j], arr[k] = arr[k], arr[j]
			j++
			k--
		}
	}

	return arr
}

func countPairs(arr []int) int {
	count = 0
	mergeSort(arr)
	return count
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr)/2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	arr = merge(left, right)

	return arr
}

func merge(left []int, right []int) []int {
	m := len(left)
	n := len(right)
	res := make([]int, m + n)
	l, r, index := 0, 0, 0

	for l < m && r < n {
		if left[l] > right[r] {
			res[index] = right[r]
			r++
			count += m - l
		} else {
			res[index] = left[l]
			l++
		}
		index++
	}

	for l < m {
		res[index] = left[l]
		index++
		l++
	}

	for r < n {
		res[index] = right[r]
		index++
		r++
	}

	return res
}
