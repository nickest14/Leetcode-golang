// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/
// Level: Medium

package leetcode

func maxMoves(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	for c := col - 2; c >= 0; c-- {
		for r := 0; r < row; r++ {
			if r > 0 && grid[r][c] < grid[r-1][c+1] {
				dp[r][c] = max(dp[r][c], dp[r-1][c+1]+1)
			}
			if grid[r][c] < grid[r][c+1] {
				dp[r][c] = max(dp[r][c], dp[r][c+1]+1)
			}
			if r < row-1 && grid[r][c] < grid[r+1][c+1] {
				dp[r][c] = max(dp[r][c], dp[r+1][c+1]+1)
			}
		}
	}

	ans := 0
	for r := 0; r < row; r++ {
		ans = max(ans, dp[r][0])
	}
	return ans
}
