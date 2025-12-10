package Leetcode

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(preorder) == 0 {
		return nil

	}
	postorderMap := make(map[int]int)
	for i, val := range postorder {
		postorderMap[val] = i
	}
	return build(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1, postorderMap)
}

func build(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int, postorderMap map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	//
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	leftrootVal := preorder[preStart+1]
	index := postorderMap[leftrootVal]
	lenSize := index - postStart + 1
	root.Left = build(preorder, preStart+1, preStart+lenSize, postorder, postStart, index, postorderMap)
	root.Right = build(preorder, preStart+lenSize+1, preEnd, postorder, index+1, postEnd-1, postorderMap)
	return root
}
