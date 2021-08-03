package Sort

// 类似347 -> 两点加强
// 1. 要存string,故无现成的结构,需要定义新的结构体(注:不要使用map,不好操作)
// 2. 对输出结果有要求->根据要求输出相应的结果,注意比较时的语句更新
func topKFrequent692(words []string, k int) []string {
	store := make(map[string]int, 0)
	for i := 0;  i < len(words); i++ {
		if _,ok := store[words[i]]; ok {
			store[words[i]]++
		} else {
			store[words[i]] = 1
		}
	}

	h := make(Heap692, 1)
	for key, v := range store { //注意，此处不能再用k了，重名了！
		cur := kv692{
			key,
			v,
		}
		h.Insert(cur)
		//fmt.Println(h.Len())
		if h.Len() > k + 1 {
			h.Pop()
		}
	}

	res := make([]string, k)
	for i := k - 1; i >= 0; i-- { // 注意题目有限定,要按单词出现频率由高到低排序, 且若不同的单词有相同出现频率,按字母顺序排序
		res[i] = (h.Pop()).string
	}
	//fmt.Println(res)
	return res
}


type kv692 struct{
	string
	int
}
type Heap692 []kv692

func (h Heap692) Len() int {
	return len(h)
}

func (h *Heap692) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap692) Insert(key kv692) {
	*h = append(*h, key)
	h.up(h.Len() - 1)
}

func (h *Heap692) up(i int) {
	parent := i/2
	if parent >= 1 && (((*h)[parent].int > (*h)[i].int) || ((*h)[parent].int == (*h)[i].int && (*h)[parent].string < (*h)[i].string)) {
		h.swap(parent, i)
		h.up(parent)
	}
}

func (h *Heap692) Pop() kv692 {
	ans := (*h)[1]
	h.swap(1, h.Len() - 1)
	*h = (*h)[: h.Len() - 1]
	h.down(1)
	return ans
}

func (h *Heap692) down(i int) {
	left := 2 * i
	right := 2 * i + 1
	min := i

	if left < h.Len() && (((*h)[left].int < (*h)[min].int) ||((*h)[left].int == (*h)[min].int && (*h)[left].string > (*h)[min].string)) {
		min = left
	}
	if right < h.Len() && (((*h)[right].int < (*h)[min].int) ||((*h)[right].int == (*h)[min].int && (*h)[right].string > (*h)[min].string)) {
		min = right
	}
	if min != i {
		h.swap(i, min)
		h.down(min)
	}
}
