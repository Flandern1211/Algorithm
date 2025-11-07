package Leetcode

func combine(n, k int) [][]int {
	var res [][]int
	var track []int

	//回溯算法核心函数
	var backtrack func(int)
	backtrack = func(start int) {
		if len(track) == k {
			res = append(res, append([]int(nil), track...))
		}

		for i := start; i < n; i++ {
			//选择
			track = append(track, n)
			//通过start参数控制树枝的遍历
			backtrack(i + 1)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(0)
	return res
}
