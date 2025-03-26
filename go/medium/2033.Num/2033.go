// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
// Level: Medium

package leetcode

import "sort"

func minOperations(grid [][]int, x int) int {
	rem := grid[0][0] % x
	arr := []int{}
	for _, row := range grid {
		for _, val := range row {
			if val%x != rem {
				return -1
			}
			arr = append(arr, val)
		}
	}

	sort.Ints(arr)
	mid := arr[len(arr)/2]
	ans := 0
	for _, val := range arr {
		ans += abs(val-mid) / x
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
