package Leetcode

// 思路:
// 先得到每个链表代表的数字.然后让它们相加得到和,和每次取余得到的余数放到一个新链表中
// 实现:已经知道每个数字的组成都可以表示为？*1 + ？*10 + ？*100....
// 例子:1234 = 4*1 + 3*10 + 2*100 + 1*1000,每次乘的就是10的几次方
// 缺陷：如果链表比较长，得到的数字可能会很大，两个相加的话可能会超过int64最大的边界，所以只适合小数字
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	count1 := 0
	count2 := 0
	for p1 != nil {
		index := 0
		count1 += p1.Val * (10 ^ index)
		index++
		p1 = p1.Next
	}
	for p2 != nil {
		index := 0
		count2 += p2.Val * (10 ^ index)
		index++
		p2 = p2.Next
	}
	count := count1 + count2
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	p := dummy.Next
	for p != nil {
		p.Val = count % 10
		p = p.Next
		count = count / 10
	}
	return dummy.Next
}

func addTwoNumbers02(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	//虚拟头结点
	dummy := &ListNode{-1, nil}
	p := dummy
	//记录进位
	carry := 0

	//	carry >0 主要为了保证最后一位进位的情况，如果两个链表长度相同且相加最后一位会进位则无法满足前两种情况，此时就需要carry
	//	这个进位标识来表示最后一位，不至于丢失
	for p1 != nil || p2 != nil || carry > 0 {
		//Val就是每个位数对应链表的数字相加再加上进位的1后的结果
		val := carry
		if p1 != nil {
			val += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			val += p2.Val
			p2 = p2.Next
		}

		carry = val / 10
		val = val % 10
		//构建新节点
		p.Next = &ListNode{val, nil}
		p = p.Next

	}
	return dummy.Next
}
