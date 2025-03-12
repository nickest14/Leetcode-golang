// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// Level: Medium

package leetcode

func numberOfSubstrings(s string) int {
	ans := 0
	left := 0
	freq := map[byte]int{'a': 0, 'b': 0, 'c': 0}

	for right := 0; right < len(s); right++ {
		freq[s[right]]++
		for freq['a'] > 0 && freq['b'] > 0 && freq['c'] > 0 {
			freq[s[left]]--
			left++
		}
		ans += left
	}
	return ans
}
