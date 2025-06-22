// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k/
// Level: Easy

package leetcode

import "strings"

func divideString(s string, k int, fill byte) []string {
	ans := []string{}
	n := len(s)
	cur := 0
	for cur < n {
		end := cur + k
		if end > n {
			end = n
		}
		ans = append(ans, s[cur:end])
		cur += k
	}
	last := ans[len(ans)-1]
	if len(last) < k {
		last += strings.Repeat(string(fill), k-len(last))
		ans[len(ans)-1] = last
	}
	return ans
}
