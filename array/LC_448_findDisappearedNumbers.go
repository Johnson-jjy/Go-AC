package array

// 解法一: 数字交换 -> 注意数组中值的取值范围 -> 下标要对应,否则会越界
func findDisappearedNumbers1(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i + 1 {
			temp := nums[i]
			for nums[i] != i + 1 && nums[i] != nums[temp - 1] {
				nums[i] = nums[temp - 1]
				nums[temp - 1] = temp
				temp = nums[i]
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i + 1 {
			res = append(res, i + 1)
		}
	}

	return res
}

// 原地hash: 利用数组中元素大小有界对响应元素进行划分
func findDisappearedNumbers2(nums []int) []int {
	res := make([]int, 0)
	n := len(nums)
	for _, v := range nums {
		v := (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			res = append(res, i + 1)
		}
	}

	return res
}