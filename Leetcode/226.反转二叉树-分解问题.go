package Leetcode

//分解问题

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	//先反转左右子树
	left := invertTree1(root.Left)
	right := invertTree1(root.Right)

	root.Left = right
	root.Right = left

	// 和定义逻辑自恰：以 root 为根的这棵二叉树已经被翻转，返回 root
	return root
}

// traverse
