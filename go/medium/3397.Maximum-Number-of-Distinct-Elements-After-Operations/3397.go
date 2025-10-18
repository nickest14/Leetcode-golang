// https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/
// Level: Medium

package leetcode

import "sort"

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	ans := 0
	prev := -(1 << 31)
	for _, num := range nums {
		low := num - k
		high := num + k
		cur := prev + 1
		if cur < low {
			cur = low
		}
		if cur <= high {
			ans++
			prev = cur
		}
	}
	return ans
}
