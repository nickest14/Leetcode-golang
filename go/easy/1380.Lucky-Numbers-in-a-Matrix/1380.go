// https://leetcode.com/problems/lucky-numbers-in-a-matrix/
// Level: Easy

package leetcode

func luckyNumbers(matrix [][]int) []int {
	minRow := make([]int, len(matrix))
	for i, row := range matrix {
		minRow[i] = row[0]
		for _, num := range row {
			if num < minRow[i] {
				minRow[i] = num
			}
		}
	}
	maxCol := make([]int, len(matrix[0]))
	for j := range maxCol {
		maxCol[j] = matrix[0][j]
		for i := range matrix {
			if matrix[i][j] > maxCol[j] {
				maxCol[j] = matrix[i][j]
			}
		}
	}

	lucky := []int{}
	for _, val := range minRow {
		for _, maxVal := range maxCol {
			if val == maxVal {
				lucky = append(lucky, val)
			}
		}
	}

	return lucky
}
