// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Level: Easy

package leetcode

func countConsistentStrings(allowed string, words []string) int {
	kv := make(map[string]bool)
	for _, v := range allowed {
		kv[string(v)] = true
	}

	count := 0
	for _, word := range words {
		for _, c := range word {
			if !kv[string(c)] {
				count++
				break
			}
		}
	}

	return len(words) - count
}
