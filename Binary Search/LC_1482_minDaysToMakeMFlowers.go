package Binary_Search

func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m * k {
		return -1
	}

	left := 1
	right := 1000000000

	ans := right
	for left <= right {
		mid := left + (right - left) / 2
		//fmt.Printf("%d\n", mid)

		if satisfyM(bloomDay, mid, m, k) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func satisfyM(arr []int, open int, m int, k int) bool {
	count := 0

	curK := 0
	for i := 0; i < len(arr); i++{
		if arr[i] <= open {
			curK++
			if curK == k {
				count++
				curK = 0
			}
		} else {
			curK = 0
		}
	}
	//fmt.Printf("count:%d\n", count)
	return count >= m
}
