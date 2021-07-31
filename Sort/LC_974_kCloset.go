package Sort

// 思路一: 快排 -> 可以直接把二维数组传入一般的quickSort函数,在其中对其进行运算即可 -> 堆排待补充
// 注: 若用hashmap存值,则无法处理距离相等的不同点的情况
var target int

func kClosest(points [][]int, k int) [][]int {
	//store := make(map[int][]int, len(points))
	//arr := make([]int, len(points))
	target = k - 1

	//for i := 0; i < len(points); i++ {
	//    distance := points[i][0] * points[i][0] + points[i][1] * points[i][1]
	//    store[distance] = points[i]
	//    arr[i] = distance
	//}

	quickSort(points, 0, len(points) - 1)
	//res := make([][]int, k)
	//for i := 0; i < k; i++ {
	//    res[i] = store[arr[i]]
	//}

	return points[:k]
}

func quickSort(arr [][]int, left int, right int) {
	if left < right {
		p := partition(arr, left, right)
		if p == target {
			return
		} else if p > target {
			quickSort(arr, left, p - 1)
		} else {
			quickSort(arr, p + 1, right)
		}
	}
}

func partition(arr [][]int, left int, right int) int {
	pivot := distance(arr[left])
	i := left + 1
	j := right
	for i <= j {
		for i <= j && distance(arr[i]) <= pivot {
			i++
		}
		for i <= j && distance(arr[j]) > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j],arr[i]
		}
	}
	arr[j], arr[left] = arr[left], arr[j]
	return j
}

func distance(arr []int) int {
	return arr[0] * arr[0] + arr[1] * arr[1]
}
