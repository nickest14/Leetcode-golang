// https://leetcode.com/problems/maximum-unique-subarray-sum-after-deletion/
// Level: Easy

package leetcode

import "math"

func maxSum(nums []int) int {
	minVal := math.MinInt32
	seen := make(map[int]bool)
	ans := 0

	for _, val := range nums {
		if !seen[val] {
			if val >= 0 {
				ans += val
			} else {
				if val > minVal {
					minVal = val
				}
			}
		}

		seen[val] = true
	}

	if ans == 0 && !seen[0] {
		return minVal
	}

	return ans
}
