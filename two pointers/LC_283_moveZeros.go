package two_pointers

// 移动零
// 思路1: 一个count对0进行计数; 另一个指针不断移动; 当前非零数时的count即为需要前移的位数
// 复杂度: 时间O(n) (注:其实需要多一个最后补0的复杂度); 空间(O(1))
func moveZeroes1(nums []int)  {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[i - count] = nums[i]
		}
	}

	for i := len(nums) - count; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 思路2: 双指针
// left左边为已经处理的非0序列; right左边为已经遍历的
// 复杂度: 同上, 省一个最后补0的操作
func moveZeroes2(nums []int)  {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}