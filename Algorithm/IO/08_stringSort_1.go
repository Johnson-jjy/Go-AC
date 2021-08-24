package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
	5
	c d a bb e
*/

func main() {
	var n int
	fmt.Scan(&n)
	store := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&store[i])
	}
	sort.Strings(store)
	strings.Join(store, " ")
	fmt.Println(store)
}
