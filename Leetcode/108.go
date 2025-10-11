package Leetcode

//convert sorted_array to binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(num []int) *TreeNode {
	if len(num) == 0 {
		return nil
	}

	mid := len(num) / 2
	head := &TreeNode{
		Val:   num[mid],
		Left:  nil,
		Right: nil,
	}
	p1, p2 := head, head
	for i := mid - 1; i >= 0; i-- {
		p1.Left = &TreeNode{
			Val:   num[i],
			Left:  nil,
			Right: nil,
		}
		p1 = p1.Left
	}
	for i := len(num) - 1; i > mid; i-- {
		p2.Right = &TreeNode{
			Val:   num[i],
			Left:  nil,
			Right: nil,
		}
		p2 = p2.Left
	}

	return head
}
