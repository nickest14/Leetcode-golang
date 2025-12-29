// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
// Level: Easy

package leetcode

func countNegatives(grid [][]int) int {
	ans := 0
	for _, row := range grid {
		firstNegative := len(row)
		for i := 0; i < len(row); i++ {
			if row[i] < 0 {
				firstNegative = i
				break
			}
		}
		ans += len(row) - firstNegative
	}
	return ans
}
