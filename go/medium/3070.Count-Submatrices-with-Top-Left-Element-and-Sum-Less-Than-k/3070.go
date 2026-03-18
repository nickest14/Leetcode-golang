// https://leetcode.com/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/
// Level: Medium

package leetcode

func countSubmatrices(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = grid[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
			if prefix[i+1][j+1] <= k {
				ans++
			}
		}
	}
	return ans
}
