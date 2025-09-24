// https://leetcode.com/problems/fraction-to-recurring-decimal/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := make([]string, 0)
	if (numerator < 0) != (denominator < 0) {
		result = append(result, "-")
	}

	abs := func(n int) int64 {
		if n < 0 {
			return int64(-n)
		}
		return int64(n)
	}

	num := abs(numerator)
	den := abs(denominator)

	result = append(result, strconv.FormatInt(num/den, 10))
	num %= den

	if num == 0 {
		return strings.Join(result, "")
	}

	result = append(result, ".")
	seen := make(map[int64]int)

	for num != 0 {
		if pos, exists := seen[num]; exists {
			result = append(result[:pos], append([]string{"("}, result[pos:]...)...)
			result = append(result, ")")
			break
		}
		seen[num] = len(result)
		num *= 10
		result = append(result, strconv.FormatInt(num/den, 10))
		num %= den
	}

	return strings.Join(result, "")
}
