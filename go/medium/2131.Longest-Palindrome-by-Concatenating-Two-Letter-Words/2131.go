// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
// Level: Medium

package leetcode

func longestPalindrome(words []string) int {
	wordFrequency := make(map[string]int)
	for _, word := range words {
		wordFrequency[word]++
	}

	palindromeLength := 0
	hasCenterWord := false

	for word, count := range wordFrequency {
		if word[0] == word[1] {
			palindromeLength += (count / 2) * 4
			if count%2 == 1 {
				hasCenterWord = true
			}
		} else if word[0] < word[1] {
			reverseWord := string([]byte{word[1], word[0]})
			reverseCount := wordFrequency[reverseWord]
			palindromeLength += min(count, reverseCount) * 4
		}
	}

	if hasCenterWord {
		palindromeLength += 2
	}

	return palindromeLength
}
