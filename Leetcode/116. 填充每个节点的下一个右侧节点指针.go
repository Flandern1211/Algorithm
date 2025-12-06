package Leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 遍历
func connect(root *Node) *Node {
	root.Next = nil
	traverse3(root)
	return root
}

func traverse3(root *Node) {
	//因为要在前序位置进行操作左右子节点，所以进行操作的左右子节点可能为nil，此时就不需要再操作了
	if root == nil || root.Left == nil || root.Right == nil {
		return
	}
	//前序位置操作

	//对该节点的左右子节点进行操作
	root.Left.Next = root.Right

	if root.Next != nil {
		if root.Next.Left == nil {
			root.Right.Next = nil
		} else {
			root.Right.Next = root.Next.Left
		}
	} else {
		root.Right.Next = nil
	}

	traverse3(root.Left)
	traverse3(root.Right)
}
