// https://leetcode.com/problems/apple-redistribution-into-boxes/
// Level: Easy

package leetcode

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	total := 0
	for _, a := range apple {
		total += a
	}

	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})

	ans := 0
	for _, cap := range capacity {
		total -= cap
		ans++
		if total <= 0 {
			return ans
		}
	}
	return ans
}
