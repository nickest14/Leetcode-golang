// https://leetcode.com/problems/house-robber-iv/
// Level: Medium

package leetcode

import "math"

func minCapability(nums []int, k int) int {
	canStealKHouses := func(capability int) bool {
		count := 0
		i := 0
		for i < len(nums) {
			if nums[i] <= capability {
				count++
				i += 2
			} else {
				i++
			}
		}
		return count >= k
	}

	left, right := 1, math.MaxInt32

	for left < right {
		mid := left + (right-left)/2
		if canStealKHouses(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
