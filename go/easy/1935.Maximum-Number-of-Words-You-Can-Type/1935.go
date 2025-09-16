// https://leetcode.com/problems/maximum-number-of-words-you-can-type/
// Level: Easy

package leetcode

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	broken := make(map[rune]bool)
	for _, letter := range brokenLetters {
		broken[letter] = true
	}
	words := strings.Split(text, " ")
	ans := 0
	for _, word := range words {
		valid := true
		for _, letter := range word {
			if broken[letter] {
				valid = false
				break
			}
		}
		if valid {
			ans++
		}
	}
	return ans
}
