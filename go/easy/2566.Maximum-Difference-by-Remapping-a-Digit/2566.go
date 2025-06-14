// https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/
// Level: Easy

package leetcode

import (
	"strconv"
	"strings"
)

func minMaxDifference(num int) int {
	numStr := strconv.Itoa(num)
	minCh := "0"
	for _, ch := range numStr {
		if ch != '0' {
			minCh = string(ch)
			break
		}
	}
	minStr := strings.ReplaceAll(numStr, minCh, "0")

	maxCh := "9"
	for _, ch := range numStr {
		if ch != '9' {
			maxCh = string(ch)
			break
		}
	}
	maxStr := strings.ReplaceAll(numStr, maxCh, "9")

	maxNum, _ := strconv.Atoi(maxStr)
	minNum, _ := strconv.Atoi(minStr)

	return maxNum - minNum
}
