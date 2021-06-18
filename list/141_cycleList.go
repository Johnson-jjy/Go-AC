package list

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	//快慢指针
	fast := head
	slow := head

	//注1：判断条件，先fast再Next
	//注2：fast和slow初始化相同，则for循环在后
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		//找到成环点，退出
		if fast == slow {
			return true
		}
	}
	return false
}