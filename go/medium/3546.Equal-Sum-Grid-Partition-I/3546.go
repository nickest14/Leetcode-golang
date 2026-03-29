// https://leetcode.com/problems/equal-sum-grid-partition-i/
// Level: Medium

package leetcode

func canPartitionGrid(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += grid[i][j]
		}
	}

	if total%2 != 0 {
		return false
	}
	target := total / 2

	s := 0
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			s += grid[i][j]
		}
		if s == target {
			return true
		}
	}

	s = 0
	for j := 0; j < n-1; j++ {
		for i := 0; i < m; i++ {
			s += grid[i][j]
		}
		if s == target {
			return true
		}
	}

	return false
}
