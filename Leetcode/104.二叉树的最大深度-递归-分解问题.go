package Leetcode

func maxDepth(root *TreeNode) int {
	//分解问题，节点的最大深度就等于其左右子节点中的最大深度+1

	//结束条件
	if root == nil {
		return 0
	}

	//分解成求左右子节点的最大深度
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
