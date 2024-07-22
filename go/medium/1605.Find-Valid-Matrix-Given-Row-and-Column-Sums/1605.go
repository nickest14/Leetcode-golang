// https://leetcode.com/problems/find-valid-matrix-given-row-and-column-sums/
// Level: Medium

package leetcode

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	rows := len(rowSum)
	cols := len(colSum)
	curRow := 0
	curCol := 0

	ans := make([][]int, rows)

	for i := range ans {
		ans[i] = make([]int, cols)
	}

	for curRow < rows || curCol < cols {
		if curRow >= rows {
			curCol++
			continue
		}
		if curCol >= cols {
			curRow++
			continue
		}

		value := min(rowSum[curRow], colSum[curCol])
		rowSum[curRow] -= value
		colSum[curCol] -= value
		ans[curRow][curCol] = value

		if rowSum[curRow] == 0 {
			curRow++
		}
		if colSum[curCol] == 0 {
			curCol++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
