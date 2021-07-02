package Binary_Search

func findDuplicate(nums []int) int {
	left := 1
	right := len(nums) - 1

	ans := right
	for left <= right {
		mid := left + (right - left) / 2

		//fmt.Printf("mid:%d\n", mid)
		if countLess(nums, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func countLess(arr []int, cur int) bool {
	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < cur {
			count++
		}
	}

	//fmt.Printf("count:%d\n", count)
	return count <= cur - 1
}
