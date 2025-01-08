// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/
// Level: Easy

package leetcode

import "strings"

func isPrefixAndSuffix(word1, word2 string) bool {
	return strings.HasPrefix(word2, word1) && strings.HasSuffix(word2, word1)
}

func countPrefixSuffixPairs(words []string) int {
	ans := 0
	n := len(words)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				ans++
			}
		}
	}
	return ans
}
