// https://leetcode.com/problems/reordered-power-of-2/
// Level: Medium

package leetcode

import (
	"sort"
	"strconv"
)

func reorderedPowerOf2(n int) bool {
	sortString := func(str string) string {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		return string(runes)
	}

	s := sortString(strconv.Itoa(n))
	for i := 0; i < 31; i++ {
		pow := 1 << i
		if s == sortString(strconv.Itoa(pow)) {
			return true
		}
	}
	return false
}
