// https://leetcode.com/problems/detect-cycles-in-2d-grid/
// Level: Medium

package leetcode

func containsCycle(grid [][]byte) bool {
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var detect func(x1, y1, p1, p2 int, prev byte) bool
	detect = func(x1, y1, p1, p2 int, prev byte) bool {
		visited[x1][y1] = true
		for _, dir := range dirs {
			dir1, dir2 := dir[0], dir[1]
			x := x1 + dir1
			y := y1 + dir2
			if dir1 != -p1 || dir2 != -p2 {
				if x >= 0 && x < m && y >= 0 && y < n && prev == grid[x][y] {
					if visited[x][y] || detect(x, y, dir1, dir2, prev) {
						return true
					}
				}
			}
		}
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if detect(i, j, -1, -1, grid[i][j]) {
					return true
				}
			}
		}
	}
	return false
}
