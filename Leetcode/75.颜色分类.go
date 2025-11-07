package Leetcode

func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	p := 0
	//由于p2为开区间，所以p2 <= p2
	for p <= p2 {
		if nums[p] == 0 {
			swap(nums, p0, p)
			p0++
		}
		if nums[p] == 2 {
			swap(nums, p2, p)
			p2--
		}
		if nums[p] == 1 {
			p++
		}
		if p < p0 {
			p = p0
		}
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]

}
