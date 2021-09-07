package main

import "fmt"

func main()  {
	var n, q int
	fmt.Scan(&n, &q)
	var str string
	for i := 0; i < q; i++ {
		fmt.Scan(&str)
		res := 0
		na := 0
		nb := 0
		nc := 0
		curFlag := -1
		for j := 0; j < n; j++ {
			if str[j] == 'A' {
				if curFlag != 0 {
					curFlag = 0
					na++
				}
			} else if str[j] == 'B' {
				if curFlag != 1 {
					curFlag = 1
					nb++
				}
			} else {
				if curFlag != 2 {
					curFlag = 2
					nc++
				}
			}
		}
		res = na + nb + nc - 3
		if na - 1 == 0 && str[0] != 'A'{
			res++
		}
		if nc - 1 == 0 && str[len(str) - 1] != 'C' {
			res++
		}
		if nb - 1

	}

}
