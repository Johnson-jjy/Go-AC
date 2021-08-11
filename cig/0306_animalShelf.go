package cig


// 注: dog和cat进队列都是有序号的,故不需要再开一个新的all保存最新的,只需要dequeueAny时比较dog和cat的最老者即可
//-> 理解题意是关键!末尾有标准错误解法
// 题意: 1.哪一个是种类,哪一个是编号,理清楚! 2.两者的编号有无联系,理清楚!->若无联系,可参照下面的解法
type AnimalShelf struct {
	dog []int
	cat []int
}

func Constructor6() AnimalShelf {
	d := make([]int, 0)
	c := make([]int, 0)
	return AnimalShelf {
		d,
		c,
	}
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	if animal[1] == 0 {
		this.cat = append(this.cat, animal[0])
	}
	if animal[1] == 1 {
		this.dog = append(this.dog, animal[0])
	}
}


func (this *AnimalShelf) DequeueAny() []int {
	res := []int{-1, -1}
	if len(this.dog) == 0 && len(this.cat) == 0 {
		return res
	}
	if len(this.cat) == 0 || (len(this.dog) != 0 && this.dog[0] < this.cat[0]) {
		res = this.DequeueDog()
	} else {
		res = this.DequeueCat()
	}
	return res
}


func (this *AnimalShelf) DequeueDog() []int {
	res := []int{-1, -1}
	if len(this.dog) == 0 {
		return res
	}
	res = make([]int, 2)
	res[1] = 1
	res[0] = this.dog[0]
	this.dog = this.dog[1:]
	return res
}


func (this *AnimalShelf) DequeueCat() []int {
	res := []int{-1, -1}
	if len(this.cat) == 0 {
		return res
	}
	res = make([]int, 2)
	res[1] = 0
	res[0] = this.cat[0]
	this.cat = this.cat[1:]
	return res
}


/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
//
//type AnimalShelf struct {
//	dog []int
//	cat []int
//	all []int
//	size int
//}
//
//func Constructor6() AnimalShelf {
//	d := make([]int, 0)
//	c := make([]int, 0)
//	a := make([]int, 0)
//	return AnimalShelf {
//		d,
//		c,
//		a,
//		0,
//	}
//}
//
//func (this *AnimalShelf) Enqueue(animal []int)  {
//	if animal[1] == 0 {
//		this.cat = append(this.cat, animal[0])
//	}
//	if animal[1] == 1 {
//		this.dog = append(this.dog, animal[0])
//	}
//	this.all = append(this.all, animal[1])
//	this.size++
//}
//
//
//func (this *AnimalShelf) DequeueAny() []int {
//	res := []int{-1, -1}
//	if this.size == 0 {
//		return res
//	}
//	res = make([]int, 2)
//	i := 0
//	for this.all[i] == -1 {
//		i++
//	}
//	if this.all[i] == 0 {
//		res = this.DequeueCat()
//	} else if this.all[i] == 1 {
//		res = this.DequeueDog()
//	}
//	return res
//}
//
//
//func (this *AnimalShelf) DequeueDog() []int {
//	res := []int{-1, -1}
//	if len(this.dog) == 0 {
//		return res
//	}
//	res = make([]int, 2)
//	res[1] = 1
//	res[0] = this.dog[0]
//	this.dog = this.dog[1:]
//	i := 0
//	for this.all[i] == 0 || this.all[i] == -1 {
//		i++
//	}
//	this.all[i] = -1
//	this.size--
//	return res
//}
//
//
//func (this *AnimalShelf) DequeueCat() []int {
//	res := []int{-1, -1}
//	if len(this.cat) == 0 {
//		return res
//	}
//	res = make([]int, 2)
//	res[1] = 0
//	res[0] = this.cat[0]
//	this.cat = this.cat[1:]
//	i := 0
//	for this.all[i] == 1 || this.all[i] == -1 {
//		i++
//	}
//	this.all[i] = -1
//	this.size--
//	return res
//}
//
//
///**
// * Your AnimalShelf object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Enqueue(animal);
// * param_2 := obj.DequeueAny();
// * param_3 := obj.DequeueDog();
// * param_4 := obj.DequeueCat();
// */