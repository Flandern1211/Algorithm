package Leetcode

import "strings"

// 方法一:先反转整个字符串再依次反转
// 可以实现原地反转，不需要再生成其他数据，占用内存少
func reverseWords(s string) string {
	var sb strings.Builder
	//清洗数据，把多余的空格都删掉
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' {
			sb.WriteByte(c)
		} else if sb.Len() > 0 && sb.String()[sb.Len()-1] != ' ' {
			//单词之间保留一个空格
			sb.WriteByte(' ')
		}
	}
	if sb.Len() == 0 {
		return ""
	}
	//删除末尾空格
	cleaned := sb.String()
	if cleaned[len(cleaned)-1] == ' ' {
		cleaned = cleaned[:len(cleaned)-1]
	}

	s1 := []rune(cleaned)
	reverseRune(s1, 0, len(s1)-1)
	start := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			reverseRune(s1, start, i-1)
			start = i + 1
		} else if i == len(s1)-1 {
			reverseRune(s1, start, i)
		}
	}
	return string(s1)
}

func reverseRune(r []rune, start int, end int) {
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}

// 方法二：先按空白字符切割为切片，然后反转切片
func reversWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]

	}
	return strings.Join(words, " ")
}
