package Sort

// 公用的思想:二分可快速求得每一排的实力
// 解法一: 排序: 归并排序的稳定 / 先用快排划分出前K个,再对前K个进行排序
// 解法二: 标准求前K--堆 -> 待补充
// 解法三: 位图思想进行扫米

// 解法一: 注map的key不能为切片,故此处开新切片记录 index -- count -> 这也是个常用的思想
func kWeakestRows1(mat [][]int, k int) []int {
	//store := make(map[[n]int]int, len(mat)) -> map的key!
	store := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		cur := make([]int, 2)
		cur[0] = i
		cur[1] = weak(mat[i])
		store[i] = cur
		//fmt.Printf("%v\n", store)
	}

	store = mergeSort1337(store)
	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = store[i][0]
	}

	return res
}

func mergeSort1337(arr [][]int) [][]int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr)/2
	left := mergeSort1337(arr[:mid])
	right := mergeSort1337(arr[mid:])
	arr = merge1337(left, right)

	return arr
}

func merge1337(left [][]int, right [][]int) [][]int {
	m, n := len(left), len(right)
	i, j, index := 0, 0, 0
	res := make([][]int, m + n)

	for i < m && j < n {
		if left[i][1] <= right[j][1] {
			res[index] = left[i]
			index++
			i++
		} else {
			res[index] = right[j]
			index++
			j++
		}
	}

	for i < m {
		res[index] = left[i]
		index++
		i++
	}
	for j < n {
		res[index] = right[j]
		index++
		j++
	}

	return res
}

func weak(arr []int) int {
	if arr[0] == 0 {
		return 0
	}
	m := len(arr)
	left := 0
	right := m - 1
	count := 0
	for left <= right {
		mid := left + (right - left)/2
		if arr[mid] == 1 {
			count = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return count + 1
}

// 解放三: 位图思想,扫描每一位
func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	mflag:=make([]int,m) // 记录该列是否已经放入了ans数组
	ans, ai := make([]int, m), 0
	// 一列一列地遍历
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mflag[j]==0 && mat[j][i] == 0 {
				mflag[j]=1
				ans[ai] = j
				ai++
			}
		}
	}
	//最后再看还有谁里面没有0的
	for i:=0;i<m;i++{
		if mflag[i]==0 {
			ans[ai]=i
			ai++
		}
	}

	return ans[:k]
}
