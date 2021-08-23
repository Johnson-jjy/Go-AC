package math

import "sort"

// 找出数组的最大公约数
// GCD模板题
func findGCD(nums []int) int {
	sort.Ints(nums)
	a := nums[0]
	b := nums[len(nums) - 1]
	res := gcd(a, b)
	return res
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
