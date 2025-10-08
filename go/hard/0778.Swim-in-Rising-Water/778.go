// https://leetcode.com/problems/swim-in-rising-water/description/?envType=daily-question&envId=2025-10-06
// Level: Hard

package leetcode

func swimInWater(grid [][]int) int {
	n := len(grid)
	directions := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	var dfs func(x, y, t int, visited [][]bool) bool
	dfs = func(x, y, t int, visited [][]bool) bool {
		if x == n-1 && y == n-1 {
			return true
		}
		visited[x][y] = true
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && !visited[nx][ny] && grid[nx][ny] <= t {
				if dfs(nx, ny, t, visited) {
					return true
				}
			}
		}
		return false
	}

	ans := n*n - 1
	low, high := 0, n*n-1
	for low <= high {
		mid := (low + high) / 2
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		if grid[0][0] <= mid && dfs(0, 0, mid, visited) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}
