package main

import (
	"fmt"
	"sort"
)

type str struct {
	origin string
	cur string
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]str, n)
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		newS := reverse(s)
		curStr := str{
			s,
			newS,
		}
		arr[i] = curStr
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i].cur < arr[j].cur
	})

	for i := 0; i < n; i++ {
		fmt.Println(arr[i].origin)
	}
}

func reverse(s string) string {
	cur := []byte(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if cur[j] < cur[i] {
			cur[i], cur[j] = cur[j], cur[i]
		}
		i++
		j--
	}
	return string(cur)
}


