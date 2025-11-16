// https://leetcode.com/problems/number-of-substrings-with-only-1s/
// Level: Medium

package leetcode

func numSub(s string) int {
	const mod = 1000000007
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		j := i
		for ; j < len(s) && s[j] == '1'; j++ {
		}
		val := (j - i)
		ans += (val * (val + 1) / 2)
		ans = ans % mod
		i = j
	}
	return ans
}
