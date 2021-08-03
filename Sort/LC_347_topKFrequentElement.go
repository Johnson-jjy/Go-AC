package Sort

// 思路:形如前K个 -> 优先队列 -> 注意在比较等方面,对于堆相关的基础语句的修改
// 本题可以作为以到模板题 -> 理解在需要统计某项值的出现此处/频率时,需要做的相关操作
func topKFrequent347(nums []int, k int) []int {
	store := make(map[int]int, 0)
	for i := 0;  i < len(nums); i++ {
		if _,ok := store[nums[i]]; ok {
			store[nums[i]]++
		} else {
			store[nums[i]] = 1
		}
	}

	h := make(Heap347, 1)
	for key, v := range store { //注意，此处不能再用k了，重名了！
		cur := make(kv347, 2)
		cur[0] = key
		cur[1] = v
		h.Insert(cur)
		//fmt.Println(h.Len())
		if h.Len() > k + 1 {
			h.Pop()
		}
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = (h.Pop())[0]
	}
	//fmt.Println(res)
	return res
}

type kv347 []int
type Heap347 []kv347

func (h Heap347) Len() int {
	return len(h)
}

func (h *Heap347) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap347) Insert(key kv347) {
	*h = append(*h, key)
	h.up(h.Len() - 1)
}

func (h *Heap347) up(i int) {
	parent := i/2
	if parent >= 1 && (*h)[parent][1] >= (*h)[i][1] {
		h.swap(parent, i)
		h.up(parent)
	}
}

func (h *Heap347) Pop() kv347 {
	ans := (*h)[1]
	h.swap(1, h.Len() - 1)
	*h = (*h)[: h.Len() - 1]
	h.down(1)
	return ans
}

func (h *Heap347) down(i int) {
	left := 2 * i
	right := 2 * i + 1
	min := i

	if left < h.Len() && (*h)[left][1] < (*h)[min][1] {
		min = left
	}
	if right < h.Len() && (*h)[right][1] < (*h)[min][1] {
		min = right
	}
	if min != i {
		h.swap(i, min)
		h.down(min)
	}
}


