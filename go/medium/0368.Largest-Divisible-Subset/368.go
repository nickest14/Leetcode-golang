// https://leetcode.com/problems/largest-divisible-subset/
// Level: Medium

package leetcode

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	maxIndex := 0
	maxSize := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
				if dp[i] > maxSize {
					maxSize = dp[i]
					maxIndex = i
				}
			}
		}
	}

	ans := []int{}
	num := nums[maxIndex]
	for i := maxIndex; i >= 0; i-- {
		if num%nums[i] == 0 && dp[i] == maxSize {
			ans = append(ans, nums[i])
			num = nums[i]
			maxSize--
		}
	}
	return ans
}
