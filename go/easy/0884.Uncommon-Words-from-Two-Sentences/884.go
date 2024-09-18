// https://leetcode.com/problems/uncommon-words-from-two-sentences/
// Level: Easy

package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	words := append(strings.Split(s1, " "), strings.Split(s2, " ")...)
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	var ans []string
	for word, count := range wordCount {
		if count == 1 {
			ans = append(ans, word)
		}
	}
	return ans
}
