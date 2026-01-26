// https://leetcode.com/problems/minimum-absolute-difference/
// Level: Easy

package leetcode

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	ans := [][]int{}
	minDiff := math.MaxInt
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			minDiff = diff
			ans = [][]int{{arr[i-1], arr[i]}}
		} else if diff == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}

	return ans
}
