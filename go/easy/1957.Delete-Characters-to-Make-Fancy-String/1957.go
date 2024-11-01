// https://leetcode.com/problems/delete-characters-to-make-fancy-string/
// Level: Easy

package leetcode

func makeFancyString(s string) string {
	if len(s) <= 2 {
		return s
	}

	ans := make([]byte, 0, len(s))
	ans = append(ans, s[0], s[1])

	for i := 2; i < len(s); i++ {
		if !(ans[len(ans)-1] == ans[len(ans)-2] && ans[len(ans)-1] == s[i]) {
			ans = append(ans, s[i])
		}
	}

	return string(ans)
}
