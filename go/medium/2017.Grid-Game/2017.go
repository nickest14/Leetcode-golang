// https://leetcode.com/problems/grid-game/

// Level: Medium

package leetcode

import "math"

func gridGame(grid [][]int) int64 {
	ans := math.MaxInt
	top, bottom := 0, 0
	for _, val := range grid[0] {
		top += val
	}

	for i := 0; i < len(grid[0]); i++ {
		top -= grid[0][i]
		ans = min(ans, max(top, bottom))
		bottom += grid[1][i]
	}
	return int64(ans)
}
