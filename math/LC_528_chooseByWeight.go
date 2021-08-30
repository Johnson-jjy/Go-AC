package math

import "math/rand"

// 按权重随机选择
// 单纯计算sum再相除,并不能获得想要的百分比->利用前缀和的方式
// 思路:生成随机seed作为目标,由初始数组得前缀和数组,继而利用二分查找seed,模拟实现了寻找随机数的效果
type Solution struct {
	sum int
	preSum []int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w) + 1)
	sum := 0
	for i := 1; i <= len(w); i++ {
		preSum[i] = preSum[i - 1] + w[i - 1]
	}
	sum = preSum[len(w)]
	return Solution {
		sum,
		preSum,
	}

}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.sum)
	//target := seed % this.sum
	left := 0
	right := len(this.preSum) - 1
	ans := right
	for left <= right {
		mid := left + (right - left)/2
		if this.preSum[mid] <= target {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}
	return ans
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
