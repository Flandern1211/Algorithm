package Leetcode

func kthSmallest(root *TreeNode, k int) int {

	res := 0
	rank := 0

	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		traverse(root.Left)
		rank++
		if k == rank {
			res = root.Val
			return
		}

		traverse(root.Right)
	}

	traverse(root)
	return res
}
