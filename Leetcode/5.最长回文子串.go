package Leetcode

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		} else {
			res = res
		}

		if len(res) < len(s2) {
			res = s2
		} else {
			res = res
		}
	}
	return res

}

// 寻找最长回文子串，分为中间为单个还是双个
// 在s中寻找以s[l]和s[r]为中心的回文子串
func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		//双指针，向两边展开
		l--
		r++
	}
	//此时s[l+1.....r-1]即是最长回文子串
	//因为是在l--和r++后的判断为false，即此时的l和r不相等
	return s[l+1 : r]
}
