// https://leetcode.com/problems/total-characters-in-string-after-transformations-i/
// Level: Medium

package leetcode

func lengthAfterTransformations(s string, t int) int {
	mod := 1000000007
	cnt := make([]int, 26)

	z := 25
	ans := len(s)

	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	for i := 0; i < t; i++ {
		ans = (ans + cnt[z]) % mod
		nextIndex := (z + 1) % 26
		cnt[nextIndex] = (cnt[nextIndex] + cnt[z]) % mod
		z = (z + 25) % 26
	}

	return ans
}
