// https://leetcode.com/problems/sentence-similarity-iii/
// Level: Medium

package leetcode

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")
	if len(words1) < len(words2) {
		words1, words2 = words2, words1
	}

	start, end := 0, 0
	n1, n2 := len(words1), len(words2)
	for start < n2 && words1[start] == words2[start] {
		start++
	}
	for end < n2 && words1[n1-1-end] == words2[n2-1-end] {
		end++
	}

	return start+end >= n2
}
