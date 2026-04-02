// https://leetcode.com/problems/maximum-amount-of-money-robot-can-earn/
// Level: Medium

package leetcode

func maximumAmount(coins [][]int) int {
	n := len(coins[0])
	const neg int64 = -1 << 58
	dp := make([][3]int64, n+1)
	for i := range dp {
		dp[i] = [3]int64{neg, neg, neg}
	}
	dp[1] = [3]int64{0, 0, 0}

	for _, row := range coins {
		for j, xv := range row {
			x := int64(xv)
			jp := j + 1
			dp[jp][2] = max(dp[j][2]+x, dp[jp][2]+x, dp[j][1], dp[jp][1])
			dp[jp][1] = max(dp[j][1]+x, dp[jp][1]+x, dp[j][0], dp[jp][0])
			dp[jp][0] = max(dp[j][0], dp[jp][0]) + x
		}
	}
	return int(dp[n][2])
}
