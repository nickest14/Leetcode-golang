// https://leetcode.com/problems/word-subsets/
// Level: Medium

package leetcode

func wordSubsets(words1 []string, words2 []string) []string {
	countFreq := func(word string) map[rune]int {
		freq := make(map[rune]int)
		for _, c := range word {
			freq[c]++
		}
		return freq
	}

	maxFreq := make(map[rune]int)
	for _, word := range words2 {
		freq := countFreq(word)
		for char, count := range freq {
			if maxFreq[char] < count {
				maxFreq[char] = count
			}
		}
	}

	var ans []string
	for _, word := range words1 {
		freq := countFreq(word)
		valid := true
		for char, count := range maxFreq {
			if freq[char] < count {
				valid = false
				break
			}
		}
		if valid {
			ans = append(ans, word)
		}
	}
	return ans
}
