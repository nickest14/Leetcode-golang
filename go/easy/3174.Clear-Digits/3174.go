// https://leetcode.com/problems/clear-digits/
// Level: Easy

package leetcode

import "unicode"

func clearDigits(s string) string {
	stack := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
