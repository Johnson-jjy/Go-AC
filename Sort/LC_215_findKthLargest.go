package Sort

// 经典题:此处用堆(优先队列->此处应该为一个大小为K的最小堆)求解,还可用快排求解
// 堆的另一种思路:构建一个大根堆,然后一次pop出的第K个值即为所求解
func findKthLargest(nums []int, k int) int {
	n := k + 1
	h := make(Heap, 1) // 注:数组首位设置为0,不参与相关操作
	for i := 0; i < len(nums); i++ {
		h.Insert(nums[i])
		if h.Len() > n {
			h.Pop()
		}
	}
	//fmt.Println(h)
	res := h.Pop() // 注意,堆顶保存的便是第K大

	return res
}

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Insert(key int) {
	*h = append(*h, key)
	h.up(h.Len() - 1)
}

func (h *Heap) up(i int) {
	parent := i/2
	if parent >= 1 && (*h)[parent] >= (*h)[i] {
		h.swap(parent, i)
		h.up(parent)
	}
}

func (h *Heap) Pop() int {
	ans := (*h)[1]
	h.swap(1, h.Len() - 1)
	*h = (*h)[: h.Len() - 1]
	h.down(1)
	return ans
}

func (h *Heap) down(i int) {
	left := 2 * i
	right := 2 * i + 1
	min := i

	if left < h.Len() && (*h)[left] < (*h)[min] {
		min = left
	}
	if right < h.Len() && (*h)[right] < (*h)[min] {
		min = right
	}
	if min != i {
		h.swap(i, min)
		h.down(min)
	}
}
