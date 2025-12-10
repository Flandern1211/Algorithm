//package Leetcode
//
//// 遍历前序遍历的每一个节点
//// 每遍历一个去中序遍历的数组中寻找相同的节点
//// 对中序遍历的数组中相同节点进行左右是否有还未使用过的数的判断
//// 左有数的话就取前序遍历的第二个节点视为左节点，然后继续对该节点进行左右是否有数的判断
//// 直到所有数都被使用过，无数可以再用时，进行右边节点的判断，同左节点
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 || len(inorder) == 0 {
//		return nil
//
//	}
//	inorderMap := make(map[int]int)
//	for i, val := range inorder {
//		inorderMap[val] = i
//	}
//
//	preIndex := 0
//
//	return build(preorder, inorder, 0, len(inorderMap)-1, inorderMap, &preIndex)
//
//}
//
//func build(preorder []int, inorder []int, inStart int, inEnd int, inorderMap map[int]int, preIndex *int) *TreeNode {
//	if inStart > inEnd {
//		return nil
//	}
//
//	rootVal := preorder[*preIndex]
//	*preIndex++
//
//	root := &TreeNode{Val: rootVal}
//
//	rootIndex := inorderMap[rootVal]
//
//	//
//	root.Left = build(preorder, inorder, inStart, rootIndex-1, inorderMap, preIndex)
//	root.Right = build(preorder, inorder, rootIndex+1, inEnd, inorderMap, preIndex)
//	return root
//}
