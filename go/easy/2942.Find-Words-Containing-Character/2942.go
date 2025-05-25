// https://leetcode.com/problems/find-words-containing-character/
// Level: Easy

package leetcode

func findWordsContaining(words []string, x byte) []int {
	ans := []int{}
	for i, word := range words {
		for _, char := range word {
			if byte(char) == x {
				ans = append(ans, i)
				break
			}
		}
	}
	return ans
}
