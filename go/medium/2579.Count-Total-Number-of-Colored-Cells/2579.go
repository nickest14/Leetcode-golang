// https://leetcode.com/problems/count-total-number-of-colored-cells/
// Level: Medium

package leetcode

func coloredCells(n int) int64 {
	if n == 1 {
		return 1
	}

	// (1 + (n - 1)) * (n-1) / 2
	corner := n * (n - 1) / 2
	return int64((2*n-1)*(2*n-1) - 4*corner)
}
