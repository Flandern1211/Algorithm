package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 1. 向序列末尾添加元素
func addlast(o *[]int, num int) {
	*o = append(*o, num)
}

// 2. 删除序列末尾元素
func remove(o *[]int) bool {
	if len(*o) == 0 {
		return true
	} else {
		*o = (*o)[:len(*o)-1]
		return false
	}
}

// 3. 获取指定索引的元素
func Index(o []int, index int) int {
	return o[index]
}

// 4. 在指定索引后插入元素
func addIndex(o *[]int, i int, x int) {
	o2 := (*o)[i+1:]
	o3 := make([]int, len((*o)[i+1:]))
	copy(o3, o2)
	*o = append((*o)[:i+1], x)
	*o = append(*o, o3...)
}

// 5. 升序排序
func sort1(o *[]int) {
	sort.Ints(*o)
}

// 6. 降序排序
func sort2(o *[]int) {
	sort.Sort(sort.Reverse(sort.IntSlice(*o)))
}

// 7. 获取序列长度
func Len(o []int) int {
	return len(o)
}

func main() {
	var o []int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(in, &op)

		switch op {
		case 1: // 添加元素到末尾
			var x int
			fmt.Fscan(in, &x)
			addlast(&o, x)
		case 2: // 删除末尾元素
			remove(&o)
		case 3: // 获取指定索引元素
			var idx int
			fmt.Fscan(in, &idx)
			if idx >= 0 && idx < len(o) {
				fmt.Fprintln(out, Index(o, idx))
			}
		case 4: // 在指定位置插入元素
			var idx, x int
			fmt.Fscan(in, &idx, &x)
			if idx >= 0 && idx < len(o) {
				addIndex(&o, idx, x)
			}
		case 5: // 升序排序
			sort1(&o)
		case 6: // 降序排序
			sort2(&o)
		case 7: // 输出序列长度
			fmt.Fprintln(out, Len(o))
		case 8: // 输出整个序列
			for j, val := range o {
				if j > 0 {
					fmt.Fprintf(out, " %d", val)
				} else {
					fmt.Fprintf(out, "%d", val)
				}
			}
			fmt.Fprintln(out)
		}
	}
}
