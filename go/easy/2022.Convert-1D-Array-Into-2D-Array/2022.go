// https://leetcode.com/problems/convert-1d-array-into-2d-array/
// Level: Easy

package leetcode

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	ans := make([][]int, 0, m)
	for i := 0; i < len(original); i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans
}
