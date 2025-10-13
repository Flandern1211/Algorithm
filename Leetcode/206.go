package Leetcode

//反转链表

// 思路：可以迭代遍历每一个节点，将其与前一个节点的连接反转
// 实现：
// 1.需要至少3个节点，目标节点，目标的父节点，目标的子节点
// 2.
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//先定义三个节点
	var pre, cur, nxt *ListNode
	pre, cur, nxt = nil, head, head.Next
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = nxt
		//注意边界条件。cur！= nil但nxt可能已经是nil了，所以必须要限制
		if nxt != nil {
			nxt = nxt.Next
		}

	}
	//应该返回pre节点，因为此时cur已经为空了
	//return cur
	return pre
}
