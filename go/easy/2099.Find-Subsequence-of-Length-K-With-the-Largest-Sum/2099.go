// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/
// Level: Easy

package leetcode

import "sort"

func maxSubsequence(nums []int, k int) []int {
	n := len(nums)
	vals := make([][]int, n)
	for i := 0; i < n; i++ {
		vals[i] = []int{i, nums[i]}
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i][1] > vals[j][1]
	})
	sort.Slice(vals[:k], func(i, j int) bool {
		return vals[i][0] < vals[j][0]
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = vals[i][1]
	}
	return ans
}
