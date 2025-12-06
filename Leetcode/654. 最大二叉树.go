package Leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {

	maxindex := 0
	max1 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max1 = nums[i]
			maxindex = i
		}
	}
	//bash case
	if maxindex < 0 || maxindex > len(nums)-1 {
		return nil
	}
	root := &TreeNode{
		Val:   max1,
		Left:  nil,
		Right: nil,
	}
	left := constructMaximumBinaryTree(nums[0:maxindex])
	right := constructMaximumBinaryTree(nums[maxindex+1 : len(nums)])

	root.Left = left
	root.Right = right
	return root
}
