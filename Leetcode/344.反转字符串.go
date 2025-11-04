package Leetcode

func reverseString(s []byte) {

	// 一左一右两个指针相向而行
	left, right := 0, len(s)-1
	for left < right {
		// 交换 s[left] 和 s[right]
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--

	}
}
