package offer

// 哈希表存储,遇重复则返回
func findRepeatNumber1(nums []int) int {
	store := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if store[nums[i]] {
			return nums[i]
		}
		store[nums[i]] = true
	}
	return -1
}

// 原地置换,时间空间均为O(1)
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			tmp := nums[nums[i]]
			if nums[i] == tmp {
				return nums[i]
			}
			nums[nums[i]] = nums[i] // 本题数据不会发生数组越界,可直接交换
			nums[i] = tmp
		}
	}
	return -1
}