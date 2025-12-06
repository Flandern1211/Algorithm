package Leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int

	//先排序，让相同的元素在一起
	sort.Ints(nums)
	backtrack(nums, &track, &res, 0)

	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, start int) {
	tmp := make([]int, len(*track))
	copy(tmp, *track)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i+1] {
			continue
		}
		*track = append(*track, nums[i])
		backtrack(nums, track, res, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
