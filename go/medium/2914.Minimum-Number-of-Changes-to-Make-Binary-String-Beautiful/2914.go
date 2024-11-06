// https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/
// Level: Medium

package leetcode

func minChanges(s string) int {
	ans := 0
	i := 0
	for i < len(s) {
		if s[i] != s[i+1] {
			ans++
		}
		i += 2
	}

	return ans
}
