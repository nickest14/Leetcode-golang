// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/
// Level: Easy

package leetcode

import "sort"

func minimumCost(nums []int) int {
	first := nums[0]
	sortedNums := make([]int, len(nums)-1)
	copy(sortedNums, nums[1:])
	sort.Ints(sortedNums)
	return first + sortedNums[0] + sortedNums[1]
}
