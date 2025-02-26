// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/
// Level: Medium

package leetcode

import "math"

func maxAbsoluteSum(nums []int) int {
	sum, minSum, maxSum := 0, 0, 0
	for _, num := range nums {
		sum += num
		if sum > maxSum {
			maxSum = sum
		}
		if sum < minSum {
			minSum = sum
		}
	}
	return int(math.Abs(float64(maxSum - minSum)))
}
