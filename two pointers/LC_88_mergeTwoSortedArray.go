package two_pointers

// 注意指针要移动!
func merge(nums1 []int, m int, nums2 []int, n int)  {
	left := m - 1
	right := n - 1
	index := m + n - 1
	for left >= 0 && right >= 0 {
		if nums1[left] > nums2[right] {
			nums1[index] = nums1[left]
			left--
		} else {
			nums1[index] = nums2[right]
			right--
		}
		index--
	}
	for right >= 0 {
		nums1[index] = nums2[right]
		index--
		right--
	}
}
