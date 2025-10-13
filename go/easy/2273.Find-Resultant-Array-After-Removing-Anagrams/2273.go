// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/
// Level: Easy

package leetcode

func removeAnagrams(words []string) []string {
	n := len(words)
	freq := make([]map[rune]int, n)
	for i, w := range words {
		freq[i] = make(map[rune]int)
		for _, ch := range w {
			freq[i][ch]++
		}
	}

	equalMaps := func(m1, m2 map[rune]int) bool {
		if len(m1) != len(m2) {
			return false
		}
		for k, v := range m1 {
			if m2[k] != v {
				return false
			}
		}
		return true
	}

	ans := []string{words[0]}
	for i := 1; i < n; i++ {
		if !equalMaps(freq[i], freq[i-1]) {
			ans = append(ans, words[i])
		}
	}

	return ans
}
