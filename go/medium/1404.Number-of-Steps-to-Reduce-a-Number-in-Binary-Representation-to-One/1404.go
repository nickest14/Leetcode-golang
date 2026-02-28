// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/
// Level: Medium

package leetcode

func numSteps(s string) int {
	steps, carry := 0, 0
	for i := len(s) - 1; i >= 1; i-- {
		bit := 0
		if s[i] == '1' {
			bit = 1
		}
		cur := bit + carry
		if cur&1 != 0 {
			steps += 2
			carry = 1
		} else {
			steps += 1
		}
	}
	return steps + carry
}
