package array

// 摩尔投票法: flag为多的那个元素, count记录次数, 若达标则直接退出; 若少于0了则夺旗
func majorityElement(nums []int) int {
	flag := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == flag {
			count++
		}
		if count > len(nums)/2{
			return flag
		}
		if nums[i] != flag {
			if count == 0 {
				flag = nums[i]
				count = 1
			} else {
				count--
			}
		}
	}

	return flag
}
