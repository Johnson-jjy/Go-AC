package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	var a, b int
	for {
		if _, err := fmt.Scan(&a, &b); err != io.EOF {
			fmt.Println(a + b)
		} else {
			break
		}
	}
}

func main(){
	var a, b, t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}

func main(){
	var a, b int

	for {
		fmt.Scan(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}

func main() {
	var n, cur int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum := 0
		for i := 0; i < n; i++ {
			fmt.Scan(&cur)
			sum += cur
		}
		fmt.Println(sum)
	}
}

func main(){
	var n, m, cur int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		sum := 0
		for j := 0; j < m; j++ {
			fmt.Scan(&cur)
			sum += cur
		}
		fmt.Println(sum)
	}
}

func main() {
	var n, cur int
	for {
		if _, err := fmt.Scan(&n); err != io.EOF {
			sum := 0
			for i := 0; i < n; i++ {
				fmt.Scan(&cur)
				sum += cur
			}
			fmt.Println(sum)
		} else {
			break
		}
	}
}

func main(){
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


func main(){
	var n int
	fmt.Scan(&n)
	b := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	sort.Strings(b)
	res := ""
	for i := 0; i < n - 1; i++ {
		res = res + b[i] + " "
	}
	res = res + b[n - 1]
	fmt.Println(res)
}

func main()  {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := strings.Split(input.Text(), " ")
		sort.StringSlice.Sort(str)
		fmt.Println(strings.Join(str, " "))
	}
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := strings.Split(input.Text(), ",")
		sort.StringSlice.Sort(str)
		fmt.Println(strings.Join(str, ","))
	}
}
