// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
// Level: Easy

package leetcode

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, k-1
	ans := math.MaxInt
	length := len(nums)
	for r < length {
		if nums[r]-nums[l] < ans {
			ans = nums[r] - nums[l]
		}
		l, r = l+1, r+1
	}
	return ans
}
