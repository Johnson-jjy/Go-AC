package main

import "fmt"

func main()  {
	var src, target string
	fmt.Scan(&src, &target)
	n := len(src)
	m := len(target)
	sum := 0
	//index := 0
	//i := 0
	check := make([]int, 26)
	store := make([]int, m)
	for j := 0; j < n; j++ {
		cur := int(src[j] - 'a')
		check[cur] = j + 1
	}
	for j := 0; j < m; j++ {
		cur := int(target[j] - 'a')
		if check[cur] == 0 {
			fmt.Println(-1)
			return
		}
		store[j] = check[cur]
	}
	fmt.Println(check, store)
	stack := make([]int, 1)
	stack[0] = store[0]
	for j := 1; j < m; j++ {
		if len(stack) > 0 && store[j] < stack[len(stack) - 1] {
			k := len(stack) - 1
			sum += n
			for k >= 0 && store[j] > stack[k] {
				k--
			}
			stack = stack[:k + 1]
		}
		stack = append(stack, store[j])
	}
	sum += stack[len(stack) - 1]
	fmt.Println(sum - m)
}

