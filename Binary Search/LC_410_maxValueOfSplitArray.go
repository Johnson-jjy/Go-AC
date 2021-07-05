package Binary_Search

func splitArray(nums []int, m int) int {

	left := 0
	right := 1000000000

	ans := right
	for left <= right {
		mid := left + (right - left) / 2
		//fmt.Printf("mid:%d\n", mid)
		if satisfyM2(nums, mid, m) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func satisfyM2(arr []int, cur int, M int) bool {
	count := 1

	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > cur {
			return false
		}
		// 用试探的方法
		if sum + arr[i] > cur {
			sum = 0
			i--
			count++
		} else {
			sum += arr[i]
		}
	}

	//fmt.Printf("count:%d\n", count)
	return count <= M
}
