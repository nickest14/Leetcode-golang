// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/
// Level: Hard

package leetcode

import (
	"fmt"
	"sort"
)

func maxValue(events [][]int, k int) int {
	memo := make(map[string]int)
	binarySearch := func(i int) int {
		key := fmt.Sprintf("binarySearch_%d", i)
		if val, exists := memo[key]; exists {
			return val
		}

		target := events[i][1] + 1
		left, right := 0, len(events)

		for left < right {
			mid := left + (right-left)/2
			if events[mid][0] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}

		memo[key] = left
		return left
	}

	var dp func(i, k int) int
	dp = func(i, k int) int {
		key := fmt.Sprintf("dp_%d_%d", i, k)
		if val, exists := memo[key]; exists {
			return val
		}

		if i >= len(events) || k <= 0 {
			return 0
		}

		skipValue := dp(i+1, k)
		next := binarySearch(i)
		takeValue := events[i][2] + dp(next, k-1)
		result := max(skipValue, takeValue)
		memo[key] = result
		return result
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	return dp(0, k)
}
