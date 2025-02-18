// https://leetcode.com/problems/construct-smallest-number-from-di-string/
// Level: Medium

package leetcode

import (
	"fmt"
	"strings"
)

func smallestNumber(pattern string) string {
	n := len(pattern)
	var s []string
	var stack []string

	for i := 0; i <= n; i++ {
		stack = append(stack, fmt.Sprintf("%d", i+1))

		if i == n || pattern[i] == 'I' {
			for len(stack) > 0 {
				s = append(s, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}

	return strings.Join(s, "")
}
