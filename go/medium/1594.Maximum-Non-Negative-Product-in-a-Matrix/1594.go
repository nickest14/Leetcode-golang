// https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/
// Level: Medium

package leetcode

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	mod := int64(1e9 + 7)

	maxDp := make([][]int64, m)
	minDp := make([][]int64, m)

	for i := range maxDp {
		maxDp[i] = make([]int64, n)
		minDp[i] = make([]int64, n)
	}

	maxDp[0][0] = int64(grid[0][0])
	minDp[0][0] = int64(grid[0][0])

	for i := 1; i < m; i++ {
		maxDp[i][0] = maxDp[i-1][0] * int64(grid[i][0])
		minDp[i][0] = maxDp[i][0]
	}

	for j := 1; j < n; j++ {
		maxDp[0][j] = maxDp[0][j-1] * int64(grid[0][j])
		minDp[0][j] = maxDp[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			val := int64(grid[i][j])

			a := maxDp[i-1][j] * val
			b := minDp[i-1][j] * val
			c := maxDp[i][j-1] * val
			d := minDp[i][j-1] * val

			maxDp[i][j] = max(a, b, c, d)
			minDp[i][j] = min(a, b, c, d)
		}
	}

	ans := maxDp[m-1][n-1]
	if ans < 0 {
		return -1
	}
	return int(ans % mod)
}
