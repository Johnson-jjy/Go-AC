package string

// 本质类似于字符串加法那几道题
// 注:不能将num数组转换为真实int，因为数组长度可能很长，int类型不够存
// 注:各种细节扣死 -> 1.对于数组,要判断索引不越界; 2.对于普通数字,要一直处理到最后; 3.对于进位,要处理,不能忘; 4.最后逆序
func addToArrayForm(num []int, k int) []int {
	res := make([]int, 0)
	carry := 0

	// 两个判断条件分别针对两种情况
	for i := len(num) - 1; i >= 0 || k > 0; i-- {
		n := 0
		if i >= 0 { //
			n = num[i]
		}
		cur := n + k%10 + carry
		carry = cur/10
		k /= 10
		res = append(res, cur%10)
	}
	if carry > 0 { // 不要忘了处理进位!
		res = append(res, carry)
	}

	for i := 0; i < len(res)/2; i++ { // 逆序更快
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
