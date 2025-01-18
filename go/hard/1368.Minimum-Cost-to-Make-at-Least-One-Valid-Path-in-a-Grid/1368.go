// https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/
// Level: Hard

package leetcode

import (
	"math"
)

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	minCost := make([][]int, m)
	for i := range minCost {
		minCost[i] = make([]int, n)
		for j := range minCost[i] {
			minCost[i][j] = math.MaxInt32
		}
	}
	minCost[0][0] = 0
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := []struct{ r, c int }{{0, 0}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell.r, cell.c

		for i, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			cost := 1
			if grid[r][c] == i+1 {
				cost = 0
			}
			if nr >= 0 && nr < m && nc >= 0 && nc < n && minCost[r][c]+cost < minCost[nr][nc] {
				minCost[nr][nc] = minCost[r][c] + cost
				if cost == 1 {
					queue = append(queue, struct{ r, c int }{nr, nc})
				} else {
					queue = append([]struct{ r, c int }{{nr, nc}}, queue...)
				}
			}
		}
	}
	return minCost[m-1][n-1]
}
