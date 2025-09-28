// https://leetcode.com/problems/largest-perimeter-triangle/
// Level: Easy

package leetcode

import "sort"

func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}
