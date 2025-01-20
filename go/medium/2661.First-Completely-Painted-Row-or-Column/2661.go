// https://leetcode.com/problems/first-completely-painted-row-or-column/
// Level: Medium

package leetcode

func firstCompleteIndex(arr []int, mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	positionMap := make(map[int][2]int)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			positionMap[mat[r][c]] = [2]int{r, c}
		}
	}

	rowCount := make([]int, rows)
	for i := 0; i < rows; i++ {
		rowCount[i] = cols
	}
	colCount := make([]int, cols)
	for i := 0; i < cols; i++ {
		colCount[i] = rows
	}

	for idx, val := range arr {
		pos := positionMap[val]
		row, col := pos[0], pos[1]
		rowCount[row]--
		colCount[col]--
		if rowCount[row] == 0 || colCount[col] == 0 {
			return idx
		}
	}
	return -1
}
