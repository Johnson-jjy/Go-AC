package main

import "fmt"

func main()  {
	var n, num int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		count := 0
		left := 0
		right := num
		ans := 0
		for left <= right {
			mid := left + (right - left)/2
			if 2 * (mid * mid  + mid) <= num {
				ans = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		count += ans * ans
		have := num - 2 * (ans * ans  + ans)
		if have >= 3 {
			count+= 1
			have -= 3
			canTwo := ans - 1
			if have <= canTwo * 2 {
				count += have/2
			} else {
				have -= canTwo * 2
				count += canTwo
				if have >= 3 {
					have -= 3
					count+= 1
					count += have/2
				}
			}
		}

		fmt.Println(count)
	}
}
