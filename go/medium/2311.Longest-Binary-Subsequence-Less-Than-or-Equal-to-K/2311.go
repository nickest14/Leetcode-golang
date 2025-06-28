// https://leetcode.com/problems/longest-binary-subsequence-less-than-or-equal-to-k/
// Level: Medium

package leetcode

import "strings"

func longestSubsequence(s string, k int) int {
	n := len(s)
	zeros := strings.Count(s, "0")
	ones := 0
	value := 0
	power := 1

	for i := n - 1; i >= 0; i-- {
		if s[i] == '1' {
			if value+power > k {
				continue
			}
			value += power
			ones++
		}
		power <<= 1
		if power > k {
			break
		}
	}
	return zeros + ones
}
