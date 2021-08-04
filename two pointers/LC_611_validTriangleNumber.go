package two_pointers

import "sort"

// 思路: 排序 + 双指针 -> 基于排序,我们应该想到用二分或者双指针,利用排序后的有序性减小算法复杂度,而不是再用循环或回溯
// 注意: 因为只需要计数,故完全没有必要去回溯
func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	index := 0
	for index < len(nums) && nums[index] == 0{
		index++
	}
	for i := index; i < len(nums); i++ {
		k := i // 此处k应该定义在j的for循环外,否则就是三重循环
		for j := i + 1; j < len(nums); j++ { // k定义在j外才可以像双指针一样有次序移动性
			for k + 1 < len(nums) && nums[k + 1] < nums[j] + nums[i] {
				k++
			}
			ans += k - j
		}
	}
	return
}

