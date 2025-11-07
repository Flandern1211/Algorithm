package Leetcode

// 双指针:
// 从nums的头尾开始遍历，以元素正负为分界线,两个指针的大小一一对比，然后放到新数组中
// 边界问题:全负，全正，都有，为空
func sortedSquares(nums []int) []int {
	//1.两个指针头尾
	i, j, k := 0, len(nums)-1, len(nums)-1
	//返回的新数组
	if len(nums) == 0 {
		return nil
	}
	nums1 := make([]int, len(nums))
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			nums1[k] = nums[i] * nums[i]
			i++
			k--
		} else {
			nums1[k] = nums[j] * nums[j]
			j--
			k--
		}
	}
	return nums1
}

//自检:

//为空:√
//全负:√
//全正:＜（＾－＾）＞
//都有:＜（＾－＾）＞
