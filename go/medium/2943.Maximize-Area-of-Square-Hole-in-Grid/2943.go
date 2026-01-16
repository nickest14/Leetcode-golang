// https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/
// Level: Medium

package leetcode

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	longest := func(bars []int) int {
		sort.Ints(bars)
		best, cur := 1, 1
		for i := 1; i < len(bars); i++ {
			if bars[i] == bars[i-1]+1 {
				cur++
			} else {
				cur = 1
			}
			if cur > best {
				best = cur
			}
		}
		return best + 1
	}

	h := longest(hBars)
	v := longest(vBars)
	if h < v {
		return h * h
	}
	return v * v
}
