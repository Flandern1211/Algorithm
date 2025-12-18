package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//判断yes的逻辑
	//先对两个都进行排序
	//然后对田忌最大的比齐威王第二大的，如果大于，则比较田忌第二大和齐威王最小的，如果同样大于才为yes
	//其他情况都为no
	horses1 := make([]int, 3)
	horses2 := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &horses1[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &horses2[i])
	}

	sort.SliceStable(horses1, func(i, j int) bool {
		return horses1[i] < horses1[j]
	})
	sort.SliceStable(horses2, func(i, j int) bool {
		return horses2[i] < horses2[j]
	})

	if horses2[2] > horses1[1] && horses2[1] > horses1[0] {
		fmt.Fprint(out, "Yes")
	} else {
		fmt.Fprint(out, "No")
	}
}
