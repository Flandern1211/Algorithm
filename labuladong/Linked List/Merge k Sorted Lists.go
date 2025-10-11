package Linked_List

import "container/heap"

// 双循环实现
// 即先循环数组中的所有链表，因为
func mergKLists(lists []*ListNode) *ListNode {
	//创建一个新的链表的虚拟节点dummy
	dummy := &ListNode{-1, nil}
	p := dummy
	for len(lists) != 0 {
		flag := 0
		min := lists[0]
		for i, _ := range lists {
			if min.Val > lists[i].Val {
				min = lists[i]
				flag = i
			}
		}
		temp := min
		lists[flag] = min.Next
		min.Next = nil
		p.Next = temp
		p = p.Next
	}
	return dummy
}

//优先级队列的实现:使用堆
//堆的使用必须先实现heap.Interface接口

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the item eventually y
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	//虚拟头结点
	dummy := &ListNode{Val: -1}
	p := dummy

	//优先级队列
	pq := &PriorityQueue{}
	heap.Init(pq)

	//将k个链表的头节点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}

	}

	for pq.Len() > 0 {
		//获取最小节点，接到结果链表中
		node := heap.Pop(pq).(*ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		//
		p = p.Next
	}
	return dummy.Next
}
