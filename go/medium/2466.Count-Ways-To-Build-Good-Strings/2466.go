// https://leetcode.com/problems/count-ways-to-build-good-strings/
// Level: Medium

package leetcode

func countGoodStrings(low int, high int, zero int, one int) int {
	const mod int = 1e9 + 7
	dp := make([]int, high+1)
	dp[0] = 1
	for i := min(zero, one); i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % mod
		}
	}
	ans := 0
	for i := low; i <= high; i++ {
		ans = (ans + dp[i]) % mod
	}
	return ans
}
