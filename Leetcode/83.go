package Leetcode

// 1.链表的分解法
func deleteDuplicates03(head *ListNode) *ListNode {
	dummy01 := &ListNode{Val: 101}
	dummy02 := &ListNode{Val: 101}

	p1, p2 := dummy01, dummy02
	p := head
	for p != nil {
		if p.Next != nil && p.Val == p.Next.Val {
			p2.Next = p
			p2 = p2.Next
		}
		if p.Next != nil && p.Val != p.Next.Val {
			p1.Next = p
			p1 = p1.Next
		}
		//注意把最后一个节点处理了
		if p.Next == nil {
			p1.Next = p
			p1 = p1.Next
		}
		p = p.Next
	}
	p1.Next = nil
	p2.Next = nil
	//返回的应该是非重复链表的虚拟头结点的下一个
	return dummy01.Next
}
