package two_pointers

import "sort"

// 三数之和: 双指针 + hash -> 固定first; second 和 third 从两端出发, 向中间寻找
// 注: 1. 因为题目需要排除重复数字,且不需要返回index -> 故排序, 每一次判断 该数是否与前一个数想等; 类似去重的全排列
// 2. 注意first需判>0, 而second则需判 > first + 1; 防止第一个second和first相等但漏了的情况
// 3. 切记,比较的是nums[first], nums[second], nums[third],事实上,这上面的写法也没规范!
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first - 1] {
			continue
		}
		if nums[first] > 0 {
			break
		}
		target := -nums[first]
		third := len(nums) - 1
		second := first + 1
		for second < third {
			if second - first > 1 && nums[second] == nums[second - 1] {
				second++
				continue
			}
			if nums[second] + nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				second++
				third--
			} else if nums[second] + nums[third] > target {
				third--
			} else {
				second++
			}
		}
	}

	return res
}

