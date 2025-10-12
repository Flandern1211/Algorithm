package Leetcode

func addTwoNumbers03(l1 *ListNode, l2 *ListNode) *ListNode {
	//将链表中的数都放到栈中
	var stk1 []int
	for l1 != nil {
		stk1 = append(stk1, l1.Val)
		l1 = l1.Next
	}

	var stk2 []int
	for l2 != nil {
		stk1 = append(stk2, l2.Val)
		l2 = l2.Next
	}

	//新链表的虚拟节点
	dummy := &ListNode{-1, nil}
	p := dummy
	//记录进位
	carry := 0

	for len(stk1) > 0 || len(stk2) > 0 || carry > 0 {

		val := carry
		if len(stk1) > 0 {
			val += stk1[len(stk1)-1]
			stk1 = stk1[:len(stk1)-1]
		}

		if len(stk2) > 0 {
			val += stk2[len(stk2)-1]
			stk1 = stk2[:len(stk2)-1]
		}
		carry = val / 10
		val = val % 10

		newListNode := &ListNode{val, nil}
		p.Next = newListNode
		p = p.Next
	}
	return dummy.Next
}
