// https://leetcode.com/problems/maximum-total-damage-with-spell-casting/
// Level: Medium

package leetcode

import "sort"

func maximumTotalDamage(power []int) int64 {
	cnt := make(map[int]int)
	for _, p := range power {
		cnt[p]++
	}

	var arr []int
	for k := range cnt {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	n := len(arr)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		val := arr[i] * cnt[arr[i]]
		j := i - 1
		for j >= 0 && arr[i]-arr[j] <= 2 {
			j--
		}
		prev := 0
		if i > 0 {
			prev = dp[i-1]
		}

		prevVal := 0
		if j >= 0 {
			prevVal = dp[j]
		}

		dp[i] = max(prev, val+prevVal)
	}

	return int64(dp[n-1])
}
