// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/
// Level: Medium

package leetcode

func countOfSubstrings(word string, k int) int64 {
	vowels := "aeiou"
	isVowel := func(c byte) bool {
		for i := range vowels {
			if c == vowels[i] {
				return true
			}
		}
		return false
	}

	atleastKConsonants := func(k int) int {
		vowelCount := make(map[byte]int)
		left, consonants, count := 0, 0, 0
		for right := 0; right < len(word); right++ {
			if isVowel(word[right]) {
				vowelCount[word[right]]++
			} else {
				consonants++
			}
			for consonants >= k && len(vowelCount) == 5 {
				count += len(word) - right
				if isVowel(word[left]) {
					vowelCount[word[left]]--
					if vowelCount[word[left]] == 0 {
						delete(vowelCount, word[left])
					}
				} else {
					consonants--
				}
				left++
			}
		}
		return count
	}

	return int64(atleastKConsonants(k) - atleastKConsonants(k+1))
}
