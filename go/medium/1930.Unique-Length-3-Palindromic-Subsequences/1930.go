// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
// Level: Medium

package leetcode

import "strings"

func countPalindromicSubsequence(s string) int {
	ans := 0
	uniq := make(map[rune]struct{})
	for _, c := range s {
		uniq[c] = struct{}{}
	}

	for c := range uniq {
		start := strings.Index(s, string(c))
		end := strings.LastIndexAny(s, string(c))
		if start < end {
			subStr := s[start+1 : end]
			uniqCharsInSubstr := make(map[rune]struct{})
			for _, ch := range subStr {
				uniqCharsInSubstr[ch] = struct{}{}
			}
			ans += len(uniqCharsInSubstr)
		}
	}
	return ans
}
