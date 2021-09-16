package array

// 求众数Ⅱ
// 摩尔投票法升级版: 1. 对于找任意大于 L/n 的情况, 最多可能由L组; 所以每次记录 L-1 组; 进行消消乐; 本质思想类似于,当3/n时,三个不一样的就直接消掉
// 2. 记录的时候,初始化设置为: count = 0; 从nums[0]开始扫, 在扫描过程中对于每一个数字都是单次计数(类似if-else的操作)
// 3. 注意每次count=0时, 修改flag之后, 会把count置为1
// 4. 2和3能保证 flag所存储的数据是不一致的; 然后在最后统计数据的时候,直接用if-else进行相关统计即可
func majorityElement229(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	res := make([]int, 0)
	flag1 := nums[0]
	flag2 := nums[0]
	count1 := 0
	count2 := 0

	for i := 0; i < len(nums); i++ {
		if flag1 == nums[i] {
			count1++
			continue
		}
		if flag2 == nums[i] {
			count2++
			continue
		}
		if count1 == 0 {
			flag1 = nums[i]
			count1 = 1
			continue
		}
		if count2 == 0 {
			flag2 = nums[i]
			count2 = 1
			//fmt.Println(flag2, i)
			continue
		}
		count1--
		count2--
	}

	//fmt.Println(flag1, flag2)
	count1 = 0
	count2 = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == flag1 {
			count1++
		} else if nums[i] == flag2 {
			count2++
		}
	}

	if count1 > len(nums)/3 {
		res = append(res, flag1)
	}
	if count2 > len(nums)/3 {
		res = append(res, flag2)
	}

	return res
}
