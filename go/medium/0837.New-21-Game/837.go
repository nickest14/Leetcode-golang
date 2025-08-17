// https://leetcode.com/problems/new-21-game/
// Level: Medium

package leetcode

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 || n >= k-1+maxPts {
		return 1.0
	}

	dp := make([]float64, n+1)
	dp[0] = 1.0
	windowSum, ans := 1.0, 0.0
	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i < k {
			windowSum += dp[i]
		} else {
			ans += dp[i]
		}
		if i >= maxPts {
			windowSum -= dp[i-maxPts]
		}
	}
	return ans
}
