// https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-i/
// Level: Medium

package leetcode

import "strings"

func maximumLength(s string) int {
	m := make(map[string]int)
	var subs []string
	current := strings.Builder{}
	for i, char := range s {
		if i == 0 || rune(s[i]) == rune(s[i-1]) {
			current.WriteRune(char)
		} else {
			subs = append(subs, current.String())
			current.Reset()
			current.WriteRune(char)
		}
	}
	subs = append(subs, current.String())

	for _, sub := range subs {
		m[sub] += 1
		if len(sub) > 1 {
			m[sub[1:]] += 2
		}
		if len(sub) > 2 {
			m[sub[2:]] += 3
		}
	}

	ans := -1
	for sub, count := range m {
		if count > 2 && len(sub) > ans {
			ans = len(sub)
		}
	}
	return ans
}
