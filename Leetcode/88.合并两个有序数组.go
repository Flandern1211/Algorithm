package Leetcode

//设计:nums2的长度为m+n,所以可以从nums1和nums2的尾部开始遍历,将nums1和nums2中的较大值排到nums1的最后，向前推进

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	k := m + n - 1

	for j >= 0 {
		if i < 0 || nums2[j] > nums1[i] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--

	}

}
