// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
// Level: Medium

package leetcode

func getHappyString(n int, k int) string {
	var dfs func(prefix string, k int) string
	dfs = func(prefix string, k int) string {
		if len(prefix) == n {
			return prefix
		}
		for _, c := range "abc" {
			if len(prefix) > 0 && prefix[len(prefix)-1] == byte(c) {
				continue
			}
			cnt := 1 << (n - len(prefix) - 1) // 2^(n - len(prefix) - 1)
			if cnt >= k {
				return dfs(prefix+string(c), k)
			} else {
				k -= cnt
			}
		}
		return ""
	}

	return dfs("", k)
}
