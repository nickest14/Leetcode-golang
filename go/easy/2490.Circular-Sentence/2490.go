// https://leetcode.com/problems/circular-sentence/
// Level: Easy

package leetcode

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		next := (i + 1) % len(words)
		if words[i][len(words[i])-1] != words[next][0] {
			return false
		}
	}

	return true
}
