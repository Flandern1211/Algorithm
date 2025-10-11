package Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	//存放小于x的链表的虚拟头结点
	dummy1 := &ListNode{-1, nil}
	//存放大于x的链表的虚拟头结点
	dummy2 := &ListNode{-1, nil}

	//p1,p2负责生成结果列表
	p1, p2 := dummy1, dummy2
	//p负责遍历源链表，
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		// 不能直接让 p 指针前进，必须在前进后断开前面的链接，否则到最后会导致两个新链表还有连接
		// p = p.Next
		// 断开原链表中的每个节点的 next 指针
		temp := p.Next
		p.Next = nil
		p = temp
	}
	// 连接两个链表
	p1.Next = dummy2.Next

	return dummy1.Next

}
