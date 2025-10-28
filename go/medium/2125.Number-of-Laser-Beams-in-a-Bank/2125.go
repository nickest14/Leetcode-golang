// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
// Level: Medium

package leetcode

import "strings"

func numberOfBeams(bank []string) int {
	ans := 0
	prev := 0
	for _, row := range bank {
		cnt := strings.Count(row, "1")
		if cnt > 0 {
			ans += prev * cnt
			prev = cnt
		}
	}
	return ans
}
