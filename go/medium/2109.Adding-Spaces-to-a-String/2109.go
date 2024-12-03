// https://leetcode.com/problems/adding-spaces-to-a-string/
// Level: Medium

package leetcode

import "strings"

func addSpaces(s string, spaces []int) string {
	var ans strings.Builder

	prev := 0
	for _, space := range spaces {
		ans.WriteString(s[prev:space])
		ans.WriteString(" ")
		prev = space
	}
	ans.WriteString(s[prev:])

	return ans.String()
}

// func addSpaces(s string, spaces []int) string {
// 	var ans []string
// 	prev := 0
// 	for _, space := range spaces {
// 		ans = append(ans, s[prev:space])
// 		prev = space
// 	}
// 	ans = append(ans, s[prev:])

// 	return strings.Join(ans, " ")
// }
