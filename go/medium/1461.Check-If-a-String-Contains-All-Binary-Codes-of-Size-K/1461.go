// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/
// Level: Medium

package leetcode

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}
	seen := make(map[string]bool)
	for i := 0; i <= len(s)-k; i++ {
		seen[s[i:i+k]] = true
	}
	return len(seen) == 1<<k
}
