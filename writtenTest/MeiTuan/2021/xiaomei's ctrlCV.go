package main

import (
	"fmt"
)

/*
输入：
5
1 2 3 4 5
5
2 1
2 5
1 2 3 4
2 3
2 5
输出：
     -1
     -1
     -1
     4
*/

// 1. 用copy比自己写for循环快
// 2. 减少一些没必要的if或for循环判据会更快
// 3. 全局变量和局部变量设置无区别
func main()  {
	var n int
	fmt.Scan(&n)
	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
		B[i] = -1
	}

	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var exec int
		fmt.Scan(&exec)
		if exec == 2 {
			var index int
			fmt.Scan(&index)
			fmt.Println(B[index - 1])
		} else {
			var k, x, y int
			fmt.Scan(&k, &x, &y)
			for k > 0 && x <= n && y <= n {
				B[y - 1] = A[x - 1]
				y++
				x++
				k--
			}
		}
	}
}


// 更快的写法见下: 快的原因未知
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//)
//
//var n, m int
//var t, k, x, y int
//
//type Order struct {
//	No    int
//	Price int
//}
//
//func main() {
//	in := bufio.NewReader(os.Stdin)
//	fmt.Fscan(in, &n)
//	A := make([]int, n)
//	B := make([]int, n)
//	for i := range A {
//		fmt.Fscan(in, &A[i])
//		B[i] = -1
//	}
//	fmt.Fscan(in, &m)
//	for ; m > 0; m-- {
//		fmt.Fscan(in, &t)
//		if t == 1 {
//			fmt.Fscan(in, &k, &x, &y)
//			copy(B[y-1:], A[x-1:x-1+k])
//		} else {
//			fmt.Fscan(in, &x)
//			fmt.Println(B[x-1])
//		}
//	}
//}
