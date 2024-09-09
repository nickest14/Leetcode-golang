// https://leetcode.com/problems/spiral-matrix-iv/
// Level: Medium

package leetcode

import (
	structures "leetcode/go/structures"
)

// ListNode is to define the struct
type ListNode = structures.ListNode

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			ans[i][j] = -1
		}
	}
	r, c := 0, 0
	dr, dc := 0, 1
	ans[r][c] = head.Val
	head = head.Next
	for head != nil {
		if 0 <= r+dr && r+dr < m && 0 <= c+dc && c+dc < n && ans[r+dr][c+dc] == -1 {
			r += dr
			c += dc
			ans[r][c] = head.Val
			head = head.Next
		} else {
			dr, dc = dc, -dr
		}
	}

	return ans
}
