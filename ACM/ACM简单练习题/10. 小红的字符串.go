package main

import (
	"fmt"
)

// 判断s是否等于t，如果等于，则需要每个字符都变化，不需要考虑最小的变化次数
//不等的话就要将s与t进行多次比对，从开头到无法往后继续推进为止，记录并一直更新一个最小次数

// 注意：纪录一个最小次数
// 替换某个字符的时候判断是该向前还是向后替换
var letters = make([]string, 26)

// 添加max和min函数定义
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Count(dst string, src string) int {
	var dstindex int
	var srcindex int
	for i, _ := range letters {
		if letters[i] == dst {
			dstindex = i
		}
		if letters[i] == src {
			srcindex = i
		}
	}
	count1 := max(dstindex, srcindex) - min(dstindex, srcindex)
	count2 := min(count1, 24-count1)
	return count2
}
func main() {
	for i := 0; i < 26; i++ {
		letters[i] = string('a' + i)
	}
	var s string
	var t string
	var mincount int = 1<<31 - 1
	fmt.Scan(&s)
	fmt.Scan(&t)

	tlen := len(t)
	slen := len(s)
	for i := 0; i < tlen-slen+1; i++ {
		count := 0
		for j := 0; j < slen; j++ {
			if s[j] == t[j+i] {
				continue
			}
			count += Count(string(t[i+j]), string(s[j]))
		}
		mincount = min(mincount, count)
	}
	fmt.Println(mincount)
}
