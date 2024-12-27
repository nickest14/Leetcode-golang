// https://leetcode.com/problems/best-sightseeing-pair/
// Level: Medium

package leetcode

func maxScoreSightseeingPair(values []int) int {
	ans := 0
	dp := make([]int, len(values))
	dp[0] = values[0]

	for i := 1; i < len(values); i++ {
		dp[i] = max(dp[i-1], values[i-1]+i-1)
		ans = max(ans, dp[i]+values[i]-i)
	}
	return ans
}
