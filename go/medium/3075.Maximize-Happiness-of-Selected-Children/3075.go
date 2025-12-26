// https://leetcode.com/problems/maximize-happiness-of-selected-children/
// Level: Medium

package leetcode

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	ans := int64(0)
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val > 0 {
			ans += int64(val)
		} else {
			break
		}
	}
	return ans
}
