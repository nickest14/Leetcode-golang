// https://leetcode.com/problems/next-greater-numerically-balanced-number/
// Level: Medium

package leetcode

import (
	"strconv"
)

func nextBeautifulNumber(n int) int {
	check := func(x int) bool {
		s := strconv.Itoa(x)
		vec := make([]int, 10)
		for _, ch := range s {
			vec[ch-'0']++
		}
		for _, ch := range s {
			c := int(ch - '0')
			if c == 0 || vec[c] != c {
				return false
			}
		}
		return true
	}

	i := n + 1
	for {
		if check(i) {
			return i
		}
		i++
	}
}
