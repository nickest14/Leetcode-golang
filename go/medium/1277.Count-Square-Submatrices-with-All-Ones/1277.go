// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
// Level: Medium

package leetcode

func countSquares(matrix [][]int) int {
	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	ans := 0

	for i := 0; i < row; i++ {
		if matrix[i][0] == 1 {
			dp[i][0] = 1
			ans++
		}
	}

	for j := 1; j < col; j++ {
		if matrix[0][j] == 1 {
			dp[0][j] = 1
			ans++
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 1 {
				min := dp[i][j-1]
				if dp[i-1][j] < min {
					min = dp[i-1][j]
				}
				if dp[i-1][j-1] < min {
					min = dp[i-1][j-1]
				}
				dp[i][j] = 1 + min
			}
			ans += dp[i][j]
		}
	}
	return ans
}
