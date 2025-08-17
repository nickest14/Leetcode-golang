// https://leetcode.com/problems/maximum-69-number/
// Level: Easy

package leetcode

import (
	"strconv"
)

func maximum69Number(num int) int {
	s := []rune(strconv.Itoa(num))
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			s[i] = '9'
			break
		}
	}
	ans, _ := strconv.Atoi(string(s))
	return ans
}
