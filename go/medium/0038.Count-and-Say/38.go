// https://leetcode.com/problems/count-and-say/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	curr := "1"
	var sb strings.Builder

	for i := 2; i <= n; i++ {
		sb.Reset()

		count := 1
		prev := curr[0]

		for j := 1; j < len(curr); j++ {
			if curr[j] == prev {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(prev)
				count = 1
				prev = curr[j]
			}
		}

		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(prev)

		curr = sb.String()
	}

	return curr
}
