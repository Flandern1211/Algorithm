package Leetcode

// 遍历:必须创建新链表

var p *TreeNode
var dummy = TreeNode{-1, nil, nil}

func flatten(root *TreeNode) {
	p = &dummy
	traverse4(root)
	if root != nil {
		root.Right = dummy.Right.Right
		root.Left = nil
	}
}
func traverse4(root *TreeNode) {
	if root == nil {
		return
	}

	//在前序的位置操作
	p.Right = &TreeNode{root.Val, nil, nil}
	p = p.Right
	traverse4(root.Left)
	traverse4(root.Right)
}

// 分解问题
func flatten2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	flatten2(root.Left)
	flatten2(root.Right)
	//保留右子树
	temp := root.Right
	//将左子树加root的右边
	root.Right = root.Left
	root.Left = nil

	p := root
	for p.Right != nil {
		p = p.Right
	}
	//拼接原右子树
	p.Right = temp

	return root

}
