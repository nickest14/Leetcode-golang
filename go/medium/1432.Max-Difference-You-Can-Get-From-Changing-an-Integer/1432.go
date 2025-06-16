// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

func maxDiff(num int) int {
	replace := func(s string, x, y byte) string {
		return strings.ReplaceAll(s, string(x), string(y))
	}

	maxStr := strconv.Itoa(num)
	for _, digit := range maxStr {
		if digit != '9' {
			maxStr = replace(maxStr, byte(digit), '9')
			break
		}
	}

	minStr := strconv.Itoa(num)
	for i := 0; i < len(minStr); i++ {
		digit := minStr[i]
		if (i == 0 && digit != '1') || (i > 0 && digit != '0' && digit != minStr[0]) {
			replaceWith := byte('0')
			if i == 0 {
				replaceWith = byte('1')
			}
			minStr = replace(minStr, digit, replaceWith)
			break
		}
	}

	max, _ := strconv.Atoi(maxStr)
	min, _ := strconv.Atoi(minStr)
	return max - min
}
