// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
// Level: Medium

package leetcode

import "sort"

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range matrix {
		dp[i] = make([]int, n)
		copy(dp[i], matrix[i])
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] += dp[i-1][j]
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		row := dp[i]
		sort.Slice(row, func(a, b int) bool { return row[a] > row[b] })
		for j := 0; j < n; j++ {
			ans = max(ans, row[j]*(j+1))
		}
	}
	return ans
}
