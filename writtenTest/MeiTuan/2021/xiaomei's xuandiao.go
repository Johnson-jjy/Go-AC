package main

import (
	"fmt"
)

/*
5
1 5 3 4 2
2 3 5 4 1
5 4 1 2 3
1 2 5 4 3
1 4 5 2 3
*/


// 只用Scan的问题: 不能及时跳转; 只用input, 如果不转换,则有空格
//func main() {
//	var n int
//	fmt.Scan(&n)
//	//var cur int
//	res := make([]int, n)
//	store := make([]bool, n)
//	input := bufio.NewScanner(os.Stdin)
//	for i := 0; i < n; i++ {
//		input.Scan()
//		str := input.Text()
//		for j := 0; j < n; j++ {
//			//fmt.Scan(&cur)
//			cur := int(str[j] - '0')
//			fmt.Println(cur)
//			if !store[cur - 1] {
//				res[i] = cur
//				store[cur - 1] = true
//				//fmt.Println(i, j)
//				break
//			}
//		}
//	}
//
//	fmt.Println(res)
//}

func main()  {
	var n int
	fmt.Scan(&n)

	res := make([]int, n)
	store := make([]bool, n)

	for i := 0; i < n; i++ {
		cur := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&cur[j])
		}
		for k := 0; k < n; k++ {
			if !store[cur[k] - 1] {
				res[i] = cur[k]
				store[cur[k] - 1] = true
				break
			}
		}
	}

	//fmt.Println(res)

	for i := 0; i < n - 1; i++ {
		fmt.Print(res[i], " ")
	}
	fmt.Print(res[n - 1])
}


//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//var n int
//
//func main() {
//	in := bufio.NewReader(os.Stdin)
//	fmt.Fscan(in, &n)
//	a := make([][]int, n)
//	for i := range a {
//		a[i] = make([]int, n)
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			fmt.Fscan(in, &a[i][j])
//		}
//	}
//	b := make([]bool, n+1)
//	var ans []int
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if b[a[i][j]] == false {
//				b[a[i][j]] = true
//				ans = append(ans, a[i][j])
//				break
//			}
//		}
//	}
//	for _, x := range ans {
//		fmt.Printf("%d ", x)
//	}
//}