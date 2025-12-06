package Leetcode

func permute(nums []int) [][]int {
	//结果集
	res := [][]int{}
	//记录的路径
	track := []int{}
	//记录使用过的元素
	used := make([]bool, len(nums))
	backtrack1(nums, track, used, &res)
	return res
}

// 路径:记录在track中
// 选择列表:nums中不存在于used中的哪些元素
// 结束条件：nums中的元素全部在track中出现
func backtrack1(nums []int, track []int, used []bool, res *[][]int) {
	//触发结束条件
	if len(nums) == len(track) {
		temp := make([]int, len(nums))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		//进入下一层决策树
		backtrack1(nums, track, used, res)
		//撤销选择
		track = track[:len(track)-1]
		used[i] = false
	}
}
