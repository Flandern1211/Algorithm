package Leetcode

// 递归遍历的方法来求最大深度
func maxDepth2(root *TreeNode) int {
	//记录每个节点的深度
	depth := 0
	//记录最大深度
	res := 0

	traverse(root, &depth, &res)
	return res

}

func traverse(root *TreeNode, depth *int, res *int) {
	//结束条件
	if root == nil {
		return
	}
	*depth++
	//到达叶子节点，此时的最大深度
	if root.Left == nil && root.Right == nil {
		*res = max(*depth, *res)
	}
	//
	traverse(root.Left, depth, res)
	traverse(root.Right, depth, res)
	*depth--
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
