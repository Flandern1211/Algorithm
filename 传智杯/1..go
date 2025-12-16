package chuanzhi

import (
	"fmt"
	"sort"
)

//https://www.nowcoder.com/practice/76babf194ed64906895fc3ece56e8a9f?tpId=405&tqId=10469871&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D405

// 蛇鸟
func main() {
	var N, L int
	fmt.Scan(&N, &L)

	heights := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&heights[i])
	}

	sort.Ints(heights)

	for i, _ := range heights {
		if L >= heights[i] {
			L++
		} else {
			break
		}
	}
	fmt.Print(L)
}
