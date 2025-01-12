// https://leetcode.com/problems/construct-k-palindrome-strings/
// Level: Medium

package leetcode

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	charFreq := make(map[rune]int)
	for _, char := range s {
		charFreq[char]++
	}

	oddCount := 0
	for _, freq := range charFreq {
		if freq%2 == 1 {
			oddCount++
		}
	}

	return oddCount <= k
}
