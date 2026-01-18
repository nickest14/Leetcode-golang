// https://leetcode.com/problems/largest-magic-square/
// Level: Medium

package leetcode

func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	row := make([][]int, m)
	for i := 0; i < m; i++ {
		row[i] = make([]int, n+1)
		for j := 0; j < n; j++ {
			row[i][j+1] = row[i][j] + grid[i][j]
		}
	}

	col := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		col[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			col[i+1][j] = col[i][j] + grid[i][j]
		}
	}

	maxSize := m
	if n < m {
		maxSize = n
	}

	for k := maxSize; k > 1; k-- {
		for i := 0; i <= m-k; i++ {
			for j := 0; j <= n-k; j++ {
				target := row[i][j+k] - row[i][j]
				ok := true

				for r := i; r < i+k; r++ {
					if row[r][j+k]-row[r][j] != target {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}

				for c := j; c < j+k; c++ {
					if col[i+k][c]-col[i][c] != target {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}

				d1 := 0
				d2 := 0
				for x := 0; x < k; x++ {
					d1 += grid[i+x][j+x]
					d2 += grid[i+x][j+k-1-x]
				}

				if d1 == target && d2 == target {
					return k
				}
			}
		}
	}

	return 1
}
