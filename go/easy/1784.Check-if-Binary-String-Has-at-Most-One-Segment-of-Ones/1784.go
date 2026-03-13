// https://leetcode.com/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
// Level: Easy

package leetcode

func checkOnesSegment(s string) bool {
	zero := false

	for _, c := range s {
		if c == '0' {
			zero = true
		} else if zero {
			return false
		}
	}
	return true
}
