package Tree

// N叉树的遍历框架
type TreeNode struct {
	Val      int
	Children []*TreeNode
}

func traverseNary(root *TreeNode) {
	if root == nil {
		return
	}
	//前序位置
	for _, child := range root.Children {
		traverseNary(child)
	}
	//后序位置
}
