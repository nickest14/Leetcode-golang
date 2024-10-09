// https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
// Level: Medium

package leetcode

func minAddToMakeValid(s string) int {
	stack := []rune{}

	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}
	return len(stack)
}
