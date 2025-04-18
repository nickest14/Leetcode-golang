// https://leetcode.com/problems/maximum-swap/
// Level: Medium

package leetcode

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numStr := []rune(fmt.Sprint(num))
	n := len(numStr)
	last := make([]int, 10)
	for i := 0; i < n; i++ {
		last[numStr[i]-'0'] = i
	}

	for i := 0; i < n; i++ {
		for d := 9; d > int(numStr[i]-'0'); d-- {
			if last[d] > i {
				numStr[i], numStr[last[d]] = numStr[last[d]], numStr[i]
				ans, _ := strconv.Atoi(string(numStr))
				return ans
			}
		}
	}
	// no swap
	return num
}
