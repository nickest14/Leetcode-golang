// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/
// Level: Easy

package leetcode

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	ans := 0
	maxValue := nums[0]
	maxDiff := 0
	for i := 1; i < n; i++ {
		ans = max(ans, maxDiff*nums[i])
		maxDiff = max(maxDiff, maxValue-nums[i])
		maxValue = max(maxValue, nums[i])
	}

	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
