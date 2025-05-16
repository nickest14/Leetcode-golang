// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/
// Level: Easy

package leetcode

func getLongestSubsequence(words []string, groups []int) []string {
	ans := []string{words[0]}
	for i := 1; i < len(words); i++ {
		if groups[i] != groups[i-1] {
			ans = append(ans, words[i])
		}
	}

	return ans
}
