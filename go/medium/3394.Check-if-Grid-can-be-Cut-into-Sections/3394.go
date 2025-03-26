// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/
// Level: Medium

package leetcode

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	canCut := func(axis int) bool {
		sort.Slice(rectangles, func(i, j int) bool {
			return rectangles[i][axis] < rectangles[j][axis]
		})

		cuts := 0
		previousEnd := -1
		for _, rect := range rectangles {
			start, end := rect[axis], rect[axis+2]
			if start >= previousEnd {
				cuts++
			}
			previousEnd = max(previousEnd, end)
			if cuts >= 3 {
				return true
			}
		}

		return false
	}

	return canCut(0) || canCut(1)
}
