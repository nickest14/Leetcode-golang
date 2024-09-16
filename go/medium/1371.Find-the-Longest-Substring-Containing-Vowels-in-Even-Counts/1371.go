// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/
// Level: Medium

package leetcode

func findTheLongestSubstring(s string) int {
	vowels := map[byte]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
	stateMap := map[int]int{0: -1}
	value := 0
	ans := 0

	for i := 0; i < len(s); i++ {
		c := s[i]
		if v, ok := vowels[c]; ok {
			value ^= v
		}
		if _, ok := stateMap[value]; !ok {
			stateMap[value] = i
		} else {
			ans = max(ans, i-stateMap[value])
		}
	}
	return ans
}
