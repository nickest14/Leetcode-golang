// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
// Level: Medium

package leetcode

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	subArraySums := make([]int, 0)
	for i := 0; i < n; i++ {
		curSum := 0
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			subArraySums = append(subArraySums, curSum)
		}
	}
	sort.Ints(subArraySums)
	ans := 0
	for i := left - 1; i < right; i++ {
		ans += subArraySums[i]
	}

	var mod = int(1e9 + 7)
	return ans % mod
}
