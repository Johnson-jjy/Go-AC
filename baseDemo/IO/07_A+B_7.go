package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	1 2 3
	4 5
	0 0 0 0 0
*/

func main()  {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		store := strings.Split(input.Text(), " ")
		sum := 0
		for i := 0; i < len(store); i++ {
			num, _ := strconv.Atoi(store[i])
			sum += num
		}
		fmt.Println(sum)
	}
}
