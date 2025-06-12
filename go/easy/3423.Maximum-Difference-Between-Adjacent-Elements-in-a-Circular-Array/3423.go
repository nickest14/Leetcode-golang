// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// Level: Easy

package leetcode

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	maxDiff := 0
	for i := 0; i < n; i++ {
		diff := abs(nums[i] - nums[(i+1)%n])
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
