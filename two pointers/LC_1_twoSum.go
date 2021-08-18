package two_pointers

// 两数之和: 哈希即可
func twoSum(nums []int, target int) []int {
	store := make(map[int]int, len(nums))

	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		need := target - cur
		if index, ok := store[need]; ok {
			res[0] = i
			res[1] = index
			break
		}
		store[cur] = i
	}

	return res
}
