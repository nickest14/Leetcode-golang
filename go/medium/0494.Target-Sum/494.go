// https://leetcode.com/problems/target-sum/
// Level: Medium

package leetcode

func findTargetSumWays(nums []int, target int) int {
	dp := make(map[[2]int]int) // (index, total) -> number of different expressions

	var backtrack func(i, total int) int
	backtrack = func(i, total int) int {
		if i == len(nums) {
			if total == target {
				return 1
			}
			return 0
		}
		key := [2]int{i, total}
		if val, exists := dp[key]; exists {
			return val
		}
		dp[key] = backtrack(i+1, total+nums[i]) + backtrack(i+1, total-nums[i])
		return dp[key]
	}

	return backtrack(0, 0)
}
