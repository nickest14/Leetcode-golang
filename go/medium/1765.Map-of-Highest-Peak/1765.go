// https://leetcode.com/problems/map-of-highest-peak/

// Level: Medium

package leetcode

func highestPeak(isWater [][]int) [][]int {
	rows := len(isWater)
	cols := len(isWater[0])
	height := make([][]int, rows)
	for i := range height {
		height[i] = make([]int, cols)
		for j := range height[i] {
			height[i][j] = 1<<32 - 1
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isWater[i][j] == 1 {
				height[i][j] = 0
			} else {
				if i > 0 {
					height[i][j] = min(height[i][j], height[i-1][j]+1)
				}
				if j > 0 {
					height[i][j] = min(height[i][j], height[i][j-1]+1)
				}
			}
		}
	}
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				height[i][j] = min(height[i][j], height[i+1][j]+1)
			}
			if j < cols-1 {
				height[i][j] = min(height[i][j], height[i][j+1]+1)
			}
		}
	}

	return height
}
