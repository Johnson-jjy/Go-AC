package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
func split( str string ) []string {
	// write code here
	res := make([]string, 0)
	if len(str) == 0 {
		return res
	}

	arr := make([][]byte, 0)
	single := make([]byte, 0)
	flag := false
	for i := 0; i < len(str); i++ {
		if str[i] == '}' {
			flag = false
			tmp := make([]byte, len(single))
			copy(tmp, single)
			arr = append(arr, tmp)
			single = make([]byte, 0)
		} else if str[i] == '{' {
			flag = true
		} else if str[i] == ',' {
			continue
		} else if flag {
			single = append(single, str[i])
		} else {
			arr = append(arr, []byte{str[i]})
		}
	}

	fmt.Println(arr)

	var dfs func(cur []byte)
	dfs = func(cur []byte) {
		if len(cur) == len(arr) {
			tmp := make([]byte, len(cur))
			copy(tmp, cur)
			res = append(res, string(tmp))
			return
		}

		chars := arr[len(cur)]
		for i := 0; i < len(chars); i++ {
			cur = append(cur, chars[i])
			dfs(cur)
			cur = cur[:len(cur) - 1]
		}
	}

	start := make([]byte, 0)
	dfs(start)

	return res
}

func main()  {
	fmt.Println(split("abcd"))
}