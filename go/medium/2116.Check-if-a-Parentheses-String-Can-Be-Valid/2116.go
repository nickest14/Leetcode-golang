// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/
// Level: Medium

package leetcode

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	openCount, closeCount := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' || locked[i] == '0' {
			openCount++
		} else {
			openCount--
		}
		if openCount < 0 {
			return false
		}
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			closeCount++
		} else {
			closeCount--
		}
		if closeCount < 0 {
			return false
		}
	}

	return true
}
