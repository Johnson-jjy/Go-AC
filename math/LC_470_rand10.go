package math

func rand7() int{
	return 1
}

// 用rand7()生成rand10() ->  核心思想:(randX() - 1)*Y + randY() 可以等概率的生成[1, X * Y]范围的随机数
func rand10() int {
	for {
		num := rand7() + (rand7() - 1) * 7
		if num <= 40 {
			return (num % 10) + 1
		}
		x := num - 40
		num = rand7() + (x - 1) * 7 // 63
		if num <= 60 {
			return (num % 10) + 1
		}
		y := num - 60
		num = rand7() + (y - 1) * 7 // 21
		if num <= 20 {
			return (num % 10) + 1
		}
	}
}
