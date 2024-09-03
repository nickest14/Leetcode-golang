// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/
// Level: Easy

package leetcode

import (
	"strconv"
)

func getLucky(s string, k int) int {
	number := ""
	for _, c := range s {
		number += strconv.Itoa(int(c) - int('a') + 1)
	}

	for i := 0; i < k; i++ {
		sum := 0
		for _, c := range number {
			digit, _ := strconv.Atoi(string(c))
			sum += digit
		}
		number = strconv.Itoa(sum)
	}

	ans, _ := strconv.Atoi(number)
	return ans
}
