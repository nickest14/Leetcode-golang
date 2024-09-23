// https://leetcode.com/problems/extra-characters-in-a-string/
// Level: Medium

package leetcode

func minExtraChar(s string, dictionary []string) int {
	maxVal := len(s) + 1
	dp := make([]int, maxVal)
	for i := range dp {
		dp[i] = maxVal
	}
	dp[0] = 0
	dicSet := make(map[string]bool)
	for _, word := range dictionary {
		dicSet[word] = true
	}
	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for j := 1; j <= i; j++ {
			if dicSet[s[i-j:i]] {
				dp[i] = min(dp[i], dp[i-j])
			}
		}
	}
	return dp[len(s)]
}
