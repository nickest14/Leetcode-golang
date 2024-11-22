// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
// Level: Medium

package leetcode

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	guard := 1
	wall := 2
	for _, g := range guards {
		grid[g[0]][g[1]] = wall
	}
	for _, w := range walls {
		grid[w[0]][w[1]] = wall
	}

	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, g := range guards {
		gx, gy := g[0], g[1]
		for _, d := range directions {
			dx, dy := d[0], d[1]
			x, y := gx, gy
			for {
				x += dx
				y += dy
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == wall {
					break
				}
				grid[x][y] = guard
			}

		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans++
			}
		}
	}
	return ans
}
