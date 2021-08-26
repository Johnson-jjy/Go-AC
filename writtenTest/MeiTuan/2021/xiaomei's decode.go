package main

import (
	"fmt"
	"strings"
)

/*
10
MMATSATMMT
*/


// bufio能显著提高输入输出的速度! 须深入研究!
// string.Index等函数并未显著提高速度
func main()  {
	var n int
	fmt.Scan(&n)
	var chars string
	fmt.Scan(&chars)
	//fmt.Println(chars)
	//i := 0
	//cM := 0
	//
	//for i < n {
	//	if chars[i] == 'M' {
	//		cM++
	//	} else if cM > 0 && chars[i] == 'T' {
	//		break
	//	}
	//	i++
	//}
	//j := n - 1
	//cT := 0
	//for j >= i {
	//	if chars[j] == 'T' {
	//		cT++
	//	} else if cT > 0 && chars[j] == 'M' {
	//		break
	//	}
	//	j--
	//}
	//fmt.Println(chars[i + 1 : j])

	PM := strings.Index(chars, "M")
	PT := strings.Index(chars[PM:], "T") + PM
	ST := strings.LastIndex(chars, "T")
	SM := strings.LastIndex(chars[:ST], "M")
	fmt.Println(chars[PT+1 : SM])
}


// 更快的做法
//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"strings"
//)
//
//var n int
//var T string
//
//func main() {
//	in := bufio.NewReader(os.Stdin)
//	fmt.Fscan(in, &n, &T)
//	PM := strings.Index(T, "M")
//	PT := strings.Index(T[PM:], "T") + PM
//	ST := strings.LastIndex(T, "T")
//	SM := strings.LastIndex(T[:ST], "M")
//	fmt.Println(T[PT+1 : SM])
//}
