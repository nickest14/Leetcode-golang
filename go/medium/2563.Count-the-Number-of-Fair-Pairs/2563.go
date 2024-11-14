// https://leetcode.com/problems/count-the-number-of-fair-pairs/
// Level: Medium

package leetcode

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	ans := 0
	for i := len(nums) - 1; i >= 1; i-- {
		val := nums[i]
		left := sort.Search(i, func(j int) bool {
			return nums[j] >= lower-val
		})
		right := sort.Search(i, func(j int) bool {
			return nums[j] > upper-val
		})
		ans += right - left
	}

	return int64(ans)
}
