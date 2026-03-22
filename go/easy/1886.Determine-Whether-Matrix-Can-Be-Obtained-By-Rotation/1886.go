// https://leetcode.com/problems/determine-whether-matrix-can-be-obtained-by-rotation/
// Level: Easy

package leetcode

func findRotation(mat [][]int, target [][]int) bool {
	matrixEqual := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if len(a[i]) != len(b[i]) {
				return false
			}
			for j := range a[i] {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
		return true
	}

	rotate90 := func(mat [][]int) [][]int {
		n := len(mat)
		out := make([][]int, n)
		for i := range out {
			out[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				out[j][n-1-i] = mat[i][j]
			}
		}
		return out
	}
	for range 4 {
		if matrixEqual(mat, target) {
			return true
		}
		mat = rotate90(mat)
	}
	return false
}
