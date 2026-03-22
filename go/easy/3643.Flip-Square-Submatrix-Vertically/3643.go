// https://leetcode.com/problems/flip-square-submatrix-vertically/
// Level: Easy

package leetcode

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		mirror := x + k - 1 - (i - x)
		for j := y; j < y+k; j++ {
			grid[i][j], grid[mirror][j] = grid[mirror][j], grid[i][j]
		}
	}
	return grid
}
