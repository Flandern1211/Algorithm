package chuanzhi

import "fmt"

//思路：根据字符串的大小，将每个字符都区分大小写提取出来放到两个切片中
//如果小写字符为空则直接返回大写字符切片的长度
//从中获取为小写的字母，并将小写的字符按顺序转为大写
//1. 无法全部转成或刚好，转成多少就是多少
//2. 可以全部转成并有多余次数，则对最后一个字符一直操作，直到次数耗尽，看是否是大写来判断是否要记录

// 变化了就返回true
func change(s rune, k int) bool {
	var n = 1
	for i := 0; i < k; i++ {
		n = -n
	}
	if n == 1 {
		return false
	}
	return true
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var s string
	fmt.Scan(&s)

	var s1 = make([]rune, 0)
	var s2 = make([]rune, 0)
	//对大小写进行分类
	for _, n := range s {
		if n >= 'a' && n <= 'z' {
			s1 = append(s1, n)
		}
		if n >= 'A' && n <= 'Z' {
			s2 = append(s2, n)
		}
	}

	//只有大写字符
	if len(s1) == 0 {
		if change(s2[0], k) {
			fmt.Println(len(s2) - 1)
			return
		} else {
			fmt.Println(len(s2))
			return
		}
	}

	//大小写都存在或者只有小写
	// k <= len(s1)
	if len(s1) >= k {
		fmt.Println(k + len(s2))
		return
	} else {
		if change(s1[0], k-len(s1)) {
			fmt.Println(len(s1) - 1 + len(s2))
			return
		} else {
			fmt.Println(len(s1) + len(s2))
			return
		}
	}

}
