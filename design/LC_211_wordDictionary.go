package design

// 添加与搜索单词 -- 数据结构设计

// 错误思路: 维护一个链表数组, 每次串起来
// 错误原因: runtime error: invalid memory address or nil pointer dereference
// line24
type node struct {
	exist bool
	next []*node
}

type WordDictionary struct {
	store []*node
}


func Constructor() WordDictionary {
	store := make([]*node, 26)
	return WordDictionary {
		store,
	}
}


func (this *WordDictionary) AddWord(word string)  {
	cur := this.store
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			if !cur[index].exist {
				cur[index].exist = true
				cur[index].next = make([]*node, 26)
			}
			cur = cur[index].next
		} else {
			for j := 0; j < 26; j++ {
				if !cur[j].exist {
					cur[j].exist = true
					cur[j].next = make([]*node, 26)
				}
				add(word[i + 1:], cur[j].next)
			}
		}
	}
}

func add(word string, cur []*node) {
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			if !cur[index].exist {
				cur[index].exist = true
				cur[index].next = make([]*node, 26)
			}
			cur = cur[index].next
		} else {
			for j := 0; j < 26; j++ {
				if !cur[j].exist {
					cur[j].exist = true
					cur[j].next = make([]*node, 26)
				}
				add(word[i + 1:], cur[j].next)
			}
		}
	}
}

func (this *WordDictionary) Search(word string) bool {
	cur := this.store
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			if !cur[index].exist {
				return false
			}
			cur = cur[index].next
		} else {
			for j := 0; j < 26; j++ {
				if cur[j].exist {
					if !search(word[i + 1:], cur) {
						return false
					}
				}
			}
		}
	}

	return true
}

func search(word string, cur []*node) bool {
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			if !cur[index].exist {
				return false
			}
			cur = cur[index].next
		} else {
			for j := 0; j < 26; j++ {
				if cur[j].exist {
					if !search(word[i + 1:], cur) {
						return false
					}
				}
			}
		}
	}
	return true
}


// 错误思路: 每次一个27位长的数组(含.对应的all)
// 错误原因: 更短的单词会被错误包含
type node struct {
	store []bool
	next *node
}

type WordDictionary struct {
	dict *node
	len int
}


func Constructor() WordDictionary {
	store := make([]bool, 27)
	dict := &node {
		store,
		nil,
	}
	return WordDictionary {
		dict,
		0,
	}
}
func (this *WordDictionary) AddWord(word string)  {
	if this.len < len(word) {
		this.len = len(word)
	}
	cur := this.dict
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			(cur.store)[index] = true
		} else {
			(cur.store)[26] = true
		}
		if cur.next == nil {
			array := make([]bool, 27)
			cur.next = &node {
				array,
				nil,
			}
		}
		cur = cur.next
	}
}


func (this *WordDictionary) Search(word string) bool {
	if len(word) > this.len {
		return false
	}
	cur := this.dict
	for i := 0; i < len(word); i++ {
		if word[i] != '.' {
			index := int(word[i] - 'a')
			if !(cur.store)[index] && !(cur.store)[26] {
				return false
			}
		}
		cur = cur.next
	}
	return true
}



/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
