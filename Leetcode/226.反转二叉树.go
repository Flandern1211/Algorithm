package Leetcode

func invertTree(root *TreeNode) *TreeNode {
	traverse2(root)
	return root
}

// traverse
func traverse2(root *TreeNode) {
	if root == nil {
		return
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	traverse2(root.Left)
	traverse2(root.Right)

}
