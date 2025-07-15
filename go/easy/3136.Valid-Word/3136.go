// https://leetcode.com/problems/valid-word/
// Level: Easy

package leetcode

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}

	hasVowel := false
	hasConsonant := false

	for _, ch := range word {
		if ch >= '0' && ch <= '9' {
			continue
		}

		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
				ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
				hasVowel = true
			} else {
				hasConsonant = true
			}
		} else {
			return false
		}
	}

	return hasVowel && hasConsonant
}
