package Leetcode

func twoSum(nums []int, target int) []int {
	// 一左一右两个指针相向而行
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			// 题目要求的索引是从 1 开始的
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 让 sum 大一点
			left++
		} else if sum > target {
			// 让 sum 小一点
			right--
		}
	}
	return []int{-1, -1}
}
