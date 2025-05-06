// https://leetcode.com/problems/domino-and-tromino-tiling/
// Level: Medium

package leetcode

func numTilings(n int) int {
	mod := int(1e9 + 7)
	size := n + 1
	if size < 4 {
		size = 4
	}
	dp := make([]int, size)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	for i := 4; i <= n; i++ {
		dp[i] = (dp[i-1]*2 + dp[i-3]) % mod
	}
	return dp[n]
}
