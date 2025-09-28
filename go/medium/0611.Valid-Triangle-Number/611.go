// https://leetcode.com/problems/valid-triangle-number/
// Level: Medium

package leetcode

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ans := 0

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				ans += right - left
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
