// https://leetcode.com/problems/rotate-string/
// Level: Easy

package leetcode

import "strings"

func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(goal+goal, s)
}
