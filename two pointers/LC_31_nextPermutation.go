package two_pointers

// 下一个排列: 找到从后往前的逆序对 -> 再扫一次从后往前,找到第一个比当前大的值(i-1)
// 注:1. 对i>0进行判断,可能全部逆序; 2.reverse的写法,可类似全部逆序,也可用i和j双向靠近
// 3.注意从i开始逆序(无论是0开始还是还是其他,都可处理) -> i+1的值已经与第一个比他大的值交换了
func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i >= 1 {
		if nums[i - 1] < nums[i] {
			break
		}
		i--
	}
	if i > 0 {
		for j := len(nums) - 1; j > i - 1; j-- {
			if nums[j] > nums[i - 1] {
				nums[i - 1], nums[j] = nums[j], nums[i - 1]
				break
			}
		}
	}
	//fmt.Println(nums)
	reverse(nums, i)
}

func reverse(nums []int, index int) {
	for i := index; i < index + (len(nums) - index)/2; i++ {
		nums[i], nums[len(nums) - i + index - 1] = nums[len(nums) - i + index - 1], nums[i]
	}
}

// for i < j {
// 	tmp := nums[i]
// 	nums[i] = nums[j]
// 	nums[j] = tmp
// 	i++
// 	j--
// }
