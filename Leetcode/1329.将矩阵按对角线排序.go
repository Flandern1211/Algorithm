package Leetcode

import "sort"

// //设计:////
// 1.根据对角巷的规则来获取每一条对角线的从左向右的索引和数值存储到两个数组中，一个为index，一个为nums
// 2.将nums数组排序后，遍历nums数组和indexs数组，将对应索引的值都变为从小到大的值

//在同一个对角线上的元素，其横纵坐标的差是相同的

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	//存储所有对角线的元素列表
	diagonals := make(map[int][]int)

	//将矩阵中的所有元素都存在map中，并按是否在同一对角线来区分
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//横纵坐标的差一定就是在同一对角线
			diagonalId := i - j
			diagonals[diagonalId] = append(diagonals[diagonalId], mat[i][j])
		}
	}

	for _, diagonal := range diagonals {
		sort.Slice(diagonal, func(i, j int) bool {
			return diagonal[i] > diagonal[j]
		})
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diagonalId := i - j
			mat[i][j] = diagonals[diagonalId][len(diagonals[diagonalId])-1]
			diagonals[diagonalId] = diagonals[diagonalId][:len(diagonals[diagonalId])-1]
		}
	}
	return mat
}
