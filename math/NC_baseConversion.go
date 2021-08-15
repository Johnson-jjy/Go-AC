package math

import (
	"fmt"
	"strconv"
)

/**
 * 进制转换
 * @param M int整型 给定整数
 * @param N int整型 转换到的进制
 * @return string字符串
 */
func solve( M int ,  N int ) string {
	// write code here
	sign := false
	if M < 0 {
		sign = true
		M = -M
	}
	var ret string
	for M > 0 {
		mod := M % N
		var a byte
		if mod <= 9 {
			a = '0' + byte(mod)
		}else{
			a = 'A' + byte(mod - 10)
		}

		ret = string(a) + ret
		M = M/N
	}
	if sign {
		ret = "-" + ret
	}
	return ret
}


func main() {
	var v int64 = 12              //默认10进制
	s2 := strconv.FormatInt(v, 2) //10 转2进制
	fmt.Printf("%v\n", s2)

	s8 := strconv.FormatInt(v, 8)
	fmt.Printf("%v\n", s8)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%v\n", s10)

	s16 := strconv.FormatInt(v, 16) //10 yo 16
	fmt.Printf("%v\n", s16)

	var sv = "11"
	fmt.Println(strconv.ParseInt(sv, 16, 32)) // 16 to 10
	fmt.Println(strconv.ParseInt(sv, 10, 32)) // 10 to 10
	fmt.Println(strconv.ParseInt(sv, 8, 32))  // 8 to 10
	fmt.Println(strconv.ParseInt(sv, 2, 32))  // 2 to 10

}