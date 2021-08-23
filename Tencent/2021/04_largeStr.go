package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var m, n int
	fmt.Scan(&m, &n)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		//fmt.Println(str)

		need := m - n
		if need == 0 {
			fmt.Println(str)
			continue
		}
		store := []byte(str)
		stack := make([]byte, 0)
		for i := 0; i < m; i++ {
			for need > 0 && len(stack) > 0 && store[i] > stack[len(stack) - 1] {
				stack = stack[:len(stack) - 1]
				need--
			}
			stack = append(stack, store[i])
		}
		fmt.Println(string(stack))
	}
}
