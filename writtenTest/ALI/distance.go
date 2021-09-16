package main

import "fmt"

/*20 2
7 2
1 7
6 8
7 9
1 2
7 0
5 4
4 4
8 9
9 5
6 6
2 6
3 1
1 1
6 4
4 0
7 1
4 0
4 3
2 1*/

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	arr := make([][]int, n)
	store1 := make(map[int][][]int, 0)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, 2)
		fmt.Scan(&arr[i][0], &arr[i][1])
		cur := store1[arr[i][0]]
		cur = append(cur, []int{arr[i][1], i})
		store1[arr[i][0]] = cur
	}

	res := 0
	//fmt.Println(store)
	//store := make([]bool, n)
	//fmt.Println(arr, k)
	for i := 0; i < n; i++ {
		minX := arr[i][0] - k
		maxX := arr[i][0] + k
		minY := arr[i][1] - k
		maxY := arr[i][1] + k
		cur1 := store[minX]
		for m := 0; m < len(cur); m++ {
				if minY <= cur[m][0] && cur[m][0] <= maxY && i < cur[m][1] {
					fmt.Println(i, cur[m][1], arr[i], arr[cur[m][1]])
					res++
				}
			}
		}
	}

	fmt.Println(res)
}

