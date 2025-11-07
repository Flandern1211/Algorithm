package Leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var track []int

	//回溯算法核心函数
	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int(nil), track...))

		for i := start; i < len(nums); i++ {
			//选择
			track = append(track, nums[i])
			//通过start参数控制树枝的遍历
			backtrack(i + 1)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
