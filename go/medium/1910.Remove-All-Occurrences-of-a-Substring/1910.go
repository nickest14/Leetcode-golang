// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/
// Level: Medium

package leetcode

func removeOccurrences(s string, part string) string {
	stack := []rune{}
	n := len(part)

	isMatch := func() bool {
		l := len(stack)
		if l < n {
			return false
		}

		for i := 0; i < n; i++ {
			if stack[l-n+i] != rune(part[i]) {
				return false
			}
		}
		return true
	}

	for _, c := range s {
		stack = append(stack, c)

		if isMatch() {
			stack = stack[:len(stack)-n]
		}
	}

	return string(stack)
}
