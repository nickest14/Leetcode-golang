// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/
// Level: Medium

package leetcode

import "math"

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			dp[i][j] = math.MaxInt

			for k := i + 1; k < j; k++ {
				score := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				if score < dp[i][j] {
					dp[i][j] = score
				}
			}
		}
	}

	return dp[0][n-1]
}
