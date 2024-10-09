// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/
// Level: Medium

package leetcode

func minSwaps(s string) int {
	maxClose, extraClose := 0, 0
	for _, c := range s {
		if c == '[' {
			extraClose--
		} else {
			extraClose++
		}
		maxClose = max(maxClose, extraClose)
	}
	return (maxClose + 1) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
