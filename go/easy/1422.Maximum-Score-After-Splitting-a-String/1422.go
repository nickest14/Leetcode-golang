// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
// Level: Easy

package leetcode

import (
	"math"
)

func maxScore(s string) int {
	zeros, ones, score := 0, 0, math.MinInt32
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '1' {
			ones++
		} else {
			zeros++
		}
		score = max(score, zeros-ones)
	}
	lastCharIsOne := 0
	if s[len(s)-1] == '1' {
		lastCharIsOne = 1
	}

	return score + ones + lastCharIsOne
}
