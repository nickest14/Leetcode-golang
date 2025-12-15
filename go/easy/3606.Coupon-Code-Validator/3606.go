// https://leetcode.com/problems/coupon-code-validator/
// Level: Easy

package leetcode

import (
	"sort"
	"unicode"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	validBusinessLines := map[string]bool{
		"electronics": true,
		"grocery":     true,
		"pharmacy":    true,
		"restaurant":  true,
	}

	groups := map[byte][]string{
		'e': {},
		'g': {},
		'p': {},
		'r': {},
	}

	isValidCode := func(code string) bool {
		for _, ch := range code {
			if !(unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_') {
				return false
			}
		}
		return true
	}

	for i := range isActive {
		if !isActive[i] || !validBusinessLines[businessLine[i]] || code[i] == "" {
			continue
		}

		if isValidCode(code[i]) {
			groups[businessLine[i][0]] = append(groups[businessLine[i][0]], code[i])
		}
	}

	order := []byte{'e', 'g', 'p', 'r'}
	var ans []string
	for _, key := range order {
		sort.Strings(groups[key])
		ans = append(ans, groups[key]...)
	}
	return ans
}
