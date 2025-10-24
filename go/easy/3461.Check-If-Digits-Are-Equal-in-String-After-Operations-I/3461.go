// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
// Level: Easy

package leetcode

func hasSameDigits(s string) bool {
	digits := make([]int, len(s))
	for i, c := range s {
		digits[i] = int(c - '0')
	}

	n := len(digits)
	for n > 2 {
		for i := 0; i < n-1; i++ {
			digits[i] = (digits[i] + digits[i+1]) % 10
		}
		n--
	}
	return digits[0] == digits[1]
}
