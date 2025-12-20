package chuanzhi

import (
	"fmt"
	"strconv"
)

//定义一个函数：专门用于将二进制转为十进制

//无限循环
//定义变量：p指针，用于将每个得到的二进制分开
//变量i 代表本次循环应该获取的字符串长度
//i = i%10
//结束条件：len(s) - p < i

func change2(s string) int64 {
	c, _ := strconv.ParseInt(s, 2, 64)
	return c
}

func main() {
	var s string
	fmt.Scan(&s)

	var s10 = make([]int64, 0)
	var p int = 0
	for i := 0; i < 100000; i++ {
		//注意用新的变量来代表需要的长度，而不是i，因为i是for里面的变量，用i会修改循环逻辑
		current := i%10 + 1
		if len(s)-p < current {
			break
		}
		s10 = append(s10, change2(s[p:p+current]))
		p = p + current
	}

	fmt.Println(len(s10))
	for i := 0; i < len(s10); i++ {
		fmt.Print(strconv.FormatInt(s10[i], 10) + " ")
	}
}
