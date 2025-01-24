// https://leetcode.com/problems/count-servers-that-communicate/

// Level: Medium

package leetcode

func countServers(grid [][]int) int {
	rows := make([]int, len(grid))
	cols := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			rows[i] += grid[i][j]
			cols[j] += grid[i][j]
		}
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
