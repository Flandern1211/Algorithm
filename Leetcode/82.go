package Leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一:根据数组可以使用索引得到前后数据的性质来实现
// 错误的做法
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	p1 := dummy
	if head == nil {
		return nil
	}

	var record []int
	var i = 2
	record[0] = head.Val
	record[1] = head.Next.Val
	cur := head.Next.Next
	for ; cur != nil; cur = cur.Next {
		record[i] = cur.Val
		if record[i-1] != record[i-2] && record[i-1] != record[i] {
			p1.Next = &ListNode{
				Val:  record[i-1],
				Next: nil,
			}
			p1 = p1.Next
		} else {
			continue
		}

	}
	return dummy
}

// 方法一的正确使用形式
func deleteDuplicatesUsingSlice(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 1) 把链表值读入 slice
	vals := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}

	n := len(vals)
	if n == 0 {
		return nil
	}

	// 2) 遍历 vals，找出“只出现一次”的值并重建链表
	dummy := &ListNode{}
	tail := dummy

	for i := 0; i < n; i++ {
		//判断当前vals[i]是否与左右不同
		leftDiff := (i == 0) || vals[i] != vals[i-1]
		rightDiff := (i == n-1) || vals[i] != vals[i+1]
		if leftDiff && rightDiff {
			tail.Next = &ListNode{Val: vals[i]}
			tail = tail.Next
		}
	}
	return dummy.Next
}

// 方法二：链表的分解
func deleteDuplicates02(head *ListNode) *ListNode {
	//将原链表分为两条
	//一条存放重复的节点，另一条都是不重复的
	dummy01 := &ListNode{Val: 101}
	dummy02 := &ListNode{Val: 101}

	p1, p2 := dummy01, dummy02
	p := head
	for p != nil {
		if p.Val == p.Next.Val || p.Val == p2.Val {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		p = p.Next
	}
	p1.Next = nil
	p2.Next = nil
	return p1.Next
}

// 快慢双指针解法
func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	p, q := dummy, head
	for q != nil {
		if q.Next != nil && q.Val == q.Next.Val {
			//发现重复数据
			for q.Next != nil && q.Val == q.Next.Val {
				q = q.Next
			}
			q = q.Next

		} else {
			p.Next = q
			q = q.Next
			p = p.Next
		}

	}
	p.Next = nil
	return dummy.Next
}

func deleteDuplicates3(head *ListNode) *ListNode {
	//将原链表分为两条
	//一条存放重复的节点，另一条都是不重复的
	dummy01 := &ListNode{Val: 101}
	dummy02 := &ListNode{Val: 101}

	p1, p2 := dummy01, dummy02
	p := head
	for p != nil {
		if p.Val == p.Next.Val || p.Val == p2.Val {
			p2.Next = p
			p2 = p2.Next
		}
		if p.Val != p.Next.Val {
			p1.Next = p
			p1 = p1.Next
		}
		p = p.Next
	}
	p1.Next = nil
	p2.Next = nil
	return p1.Next
}
