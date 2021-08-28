package two_pointers

// 两数之和Ⅱ--输入有序数组
// 数组有序,故可用双指针
func twoSum167(numbers []int, target int) []int {
	res := make([]int, 2)
	i := 0
	j := len(numbers) - 1
	for i < j {
		cur := numbers[i] + numbers[j]
		if cur < target {
			i++
		} else if cur > target {
			j--
		} else {
			res[0] = i + 1
			res[1] = j + 1
			break
		}
	}

	return res
}
