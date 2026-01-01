// https://leetcode.com/problems/plus-one/
// Level: Easy

package leetcode

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		carry, digits[i] = sum/10, sum%10
		if carry == 0 {
			return digits
		}
	}
	return append([]int{1}, digits...)
}
