// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-ii/
// Level: Medium

package leetcode

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	hammingDistance := func(s1, s2 string) int {
		count := 0
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				count++
			}
		}
		return count
	}

	n := len(words)
	dpMemo := make(map[int][]string)

	var dp func(i int) []string
	dp = func(i int) []string {
		if chain, ok := dpMemo[i]; ok {
			return chain
		}

		longestChain := []string{}
		for j := i + 1; j < n; j++ {
			if groups[i] != groups[j] && len(words[i]) == len(words[j]) && hammingDistance(words[i], words[j]) == 1 {
				candidate := dp(j)
				if len(candidate) > len(longestChain) {
					longestChain = candidate
				}
			}
		}
		result := append([]string{words[i]}, longestChain...)
		dpMemo[i] = result
		return result
	}

	var ans []string
	for i := 0; i < n; i++ {
		candidate := dp(i)
		if len(candidate) > len(ans) {
			ans = candidate
		}
	}

	return ans

}
