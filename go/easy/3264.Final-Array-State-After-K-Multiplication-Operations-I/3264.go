// https://leetcode.com/problems/final-array-state-after-k-multiplication-operations-i/
// Level: Easy

package leetcode

import "math"

func getFinalState(nums []int, k int, multiplier int) []int {
	getMinIndex := func() int {
		minVal := math.MaxInt
		minIndex := -1
		for i, num := range nums {
			if num < minVal {
				minVal = num
				minIndex = i
			}
		}
		return minIndex
	}

	for k > 0 {
		index := getMinIndex()
		nums[index] *= multiplier
		k--
	}
	return nums
}
