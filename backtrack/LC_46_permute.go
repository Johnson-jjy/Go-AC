package backtrack

// 标准全排列: 1.全局数组 2.当前数组 3.used数组 4.每次从0开始
var res [][]int

func permute(nums []int) [][]int {
	res = make([][]int, 0)

	cur := make([]int, 0)
	used := make([]int, len(nums))
	backtrack(cur, nums, used)

	return res
}

func backtrack(cur []int, nums []int, used []int) {
	if len(cur) == len(nums) {
		tmp := make([]int, len(cur))
		//fmt.Println(tmp, res)
		copy(tmp, cur)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] == 1 {
			continue
		}
		cur = append(cur, nums[i])
		used[i] = 1
		backtrack(cur, nums, used)
		used[i] = 0
		cur = cur[:len(cur) - 1]
	}
}
