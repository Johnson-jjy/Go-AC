package Binary_Search

// 搜索旋转排序数组: 互不相同 -> 可直接从left和right出发进行判断
// nums[mid]可能出现在两个位置,target可能出现在四个位置,每一个nums[mid]对应4个target位置,分类讨论即可
func search1(nums []int, target int) int {
	m := len(nums)
	left := 0
	right := m - 1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// 注意，这里nums[mid]与nums[0]相当的比较实际上是考虑了mid刚好为0时的情况
			if target < nums[0] && nums[mid] >= nums[0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// 同上，这里target与nums[0]相当的比较实际上是考虑了二者相等的情况
			// 注：这两处的本质都是考虑的如果num[0]等于target或者nums[mid]会怎样，取等只是根据各自情况所需进行选择
			if target >= nums[0] && nums[mid] < nums[0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
