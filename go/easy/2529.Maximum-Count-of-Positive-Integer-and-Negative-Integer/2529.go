// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/
// Level: Easy

package leetcode

func maximumCount(nums []int) int {
	n := len(nums)

	left, right := 0, n-1
	// Find the index of the first positive number
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	positiveCount := n - left

	left, right = 0, n-1
	// Find the last negative number using binary search
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	negativeCount := right + 1

	if positiveCount > negativeCount {
		return positiveCount
	}
	return negativeCount
}
