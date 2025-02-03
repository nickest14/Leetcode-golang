// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
// Level: Easy

package leetcode

func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	inc, dec := 1, 1
	ans := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		} else if nums[i] < nums[i-1] {
			dec++
			inc = 1
		} else {
			inc = 1
			dec = 1
		}
		ans = max(ans, inc, dec)
	}
	return ans
}
