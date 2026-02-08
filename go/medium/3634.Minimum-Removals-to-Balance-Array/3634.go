// https://leetcode.com/problems/minimum-removals-to-balance-array/
// Level: Medium

package leetcode

import "sort"

func minRemoval(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return 0
	}
	if nums[0]*k >= nums[n-1] {
		return 0
	}

	left := 0
	maxLen := 0
	for right := 0; right < n; right++ {
		for left < right && nums[left]*k < nums[right] {
			left++
		}

		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return n - maxLen
}
