// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// Level: Easy

package leetcode

import (
	"strconv"
	"strings"
)

func getNoZeroIntegers(n int) []int {
	check := func(x int) bool {
		return !strings.Contains(strconv.Itoa(x), "0")
	}

	for i := 1; i < n; i++ {
		j := n - i
		if check(i) && check(j) {
			return []int{i, j}
		}
	}
	return []int{}
}
