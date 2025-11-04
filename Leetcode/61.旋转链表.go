package Leetcode

// 1.根据k的值分隔开链表，然后连接
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算长度并找到尾
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	//2.k的值可能大于链表长度，这时其实就相当于转了几圈后又移动了几个值，也就是取余的结果
	k = k % n
	if k == 0 {
		return head
	}

	// 找到第 n-k 个节点 (firstTail)
	firstTail := head
	for i := 1; i < n-k; i++ {
		firstTail = firstTail.Next
	}
	secondHead := firstTail.Next
	firstTail.Next = nil // 断开

	// 找到 second 的尾 (secondTail)
	secondTail := secondHead
	for secondTail.Next != nil {
		secondTail = secondTail.Next
	}

	// 接上原来的 head
	secondTail.Next = head
	return secondHead
}
