// https://leetcode.com/problems/maximum-beauty-of-an-array-after-applying-operation/
// Level: Medium

package leetcode

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	left := 0
	ans := 0

	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}

		ans = max(ans, right-left+1)
	}
	return ans
}
