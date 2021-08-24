package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
	a,c,bb
	f,dddd
	nowcoder
*/

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		strs := strings.Split(input.Text(), " ")
		sort.Strings(strs)
		fmt.Println(strings.Join(strs, " "))
	}
}
