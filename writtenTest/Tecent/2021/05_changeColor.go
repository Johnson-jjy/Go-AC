package main

import (
	"fmt"
	"math"
)

var res int

func main() {
	res = math.MaxInt32

	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println(0)
		return
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	//fmt.Println("start")
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		copy(tmp, arr)
		j := i
		for j < len(tmp) && tmp[j] == tmp[i] {
			j++
		}
		j--
		//fmt.Println("?", i, tmp)
		backtrack(tmp, 0, i, j)
		i = j + 1
	}
	fmt.Println(res)
}

func backtrack(cur []int, sum int, left int, right int) bool {
	//fmt.Println("?", left, right, sum, cur)
	//time.Sleep(time.Second)
	if sum >= res {
		return false
	}
	if left == 0 && right == len(cur) - 1 {
		res = sum
		return true
	}

	if left > 0 && right < len(cur) - 1 {
		may1 := cur[left - 1]
		may2 := cur[right + 1]
		diff1 := abs(may1, cur[left])
		diff2 := abs(may2, cur[right])
		if diff1 < diff2 {
			cur[left] = may1
			cur[right] = may1
			sum += diff1 * (right - left + 1)
		} else {
			cur[left] = may2
			cur[right] = may2
			sum += diff2 * (right - left + 1)
		}
		flag := cur[left]
		for left >= 0 && right < len(cur) {
			if cur[left] == flag {
				left--
			}else {
				break
			}
			if cur[right] == flag {
				right++
			} else {
				break
			}
		}
		left++
		right--
		//fmt.Println("!1", sum, left, right, cur)
		if backtrack(cur, sum, left, right) {
			return true
		}
	} else if left == 0 {
		may := cur[right + 1]
		diff := abs(may, cur[right])
		sum += diff * (right - left + 1)
		cur[left] = may
		cur[right] = may
		for right < len(cur) {
			if cur[right] == may {
				right++
			} else {
				break
			}
		}
		right--
		//fmt.Println("!2", sum, left, right, cur)
		if backtrack(cur, sum, left, right) {
			return true
		}
	} else {
		//fmt.Println("!3", sum, left, right, cur)
		may := cur[left - 1]
		diff := abs(may, cur[left])
		sum += diff * (right - left + 1)
		cur[left] = may
		cur[right] = may
		for left >= 0 {
			if cur[left] == may {
				left--
			} else {
				break
			}
		}
		left++
		if backtrack(cur, sum, left, right) {
			return true
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//func min(a, b int) int {
//	if a > b {
//		return b
//	}
//	return a
//}