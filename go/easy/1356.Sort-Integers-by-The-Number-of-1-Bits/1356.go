// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/
// Level: Easy

package leetcode

import "sort"

func sortByBits(arr []int) []int {
	countBits := func(x int) int {
		count := 0
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return count
	}

	ans := make([]int, len(arr))
	copy(ans, arr)
	sort.Slice(ans, func(i, j int) bool {
		ci, cj := countBits(ans[i]), countBits(ans[j])
		if ci != cj {
			return ci < cj
		}
		return ans[i] < ans[j]
	})
	return ans
}
