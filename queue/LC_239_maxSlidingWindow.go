package queue

// 滑动窗口最大值
// 解法一: 单调队列: 理清楚走的是谁-->先进的对头元素; 理清楚进的是谁-->比队头元素小才进,比队头元素大则删除对头元素
func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)

	for i := 0; i < k; i++ {
		j := len(queue) - 1
		for j >= 0 && nums[i] > queue[j] {
			j--
		}
		queue = queue[ : j + 1]
		queue = append(queue, nums[i])
	}
	res := make([]int, len(nums) - k + 1)
	left := 1
	right := k
	res[0] = queue[0]
	for right < len(nums){
		if nums[left - 1] == queue[0] {
			queue = queue[1:]
		}
		j := len(queue) - 1
		for j >= 0 && nums[right] > queue[j] {
			j--
		}
		queue = queue[ : j + 1]
		queue = append(queue, nums[right])
		res[left] = queue[0]
		left++
		right++
	}

	return res
}
