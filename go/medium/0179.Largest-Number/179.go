// https://leetcode.com/problems/largest-number/
// Level: Medium

package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(strs); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[j]
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
