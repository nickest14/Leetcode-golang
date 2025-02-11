// https://leetcode.com/problems/counting-words-with-a-given-prefix/
// Level: Easy

package leetcode

import "strings"

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			ans++
		}
	}
	return ans
}
