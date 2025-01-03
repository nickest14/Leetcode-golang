// https://leetcode.com/problems/count-vowel-strings-in-ranges/
// Level: Medium

package leetcode

func vowelStrings(words []string, queries [][]int) []int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	n := len(words)
	prefix := make([]int, n+1)
	for i, word := range words {
		prefix[i+1] = prefix[i]
		if vowels[word[0]] && vowels[word[len(word)-1]] {
			prefix[i+1]++
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}
