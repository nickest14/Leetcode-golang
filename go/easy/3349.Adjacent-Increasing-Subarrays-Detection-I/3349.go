// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
// Level: Easy

package leetcode

func hasIncreasingSubarrays(nums []int, k int) bool {
	inc, prevInc, maxLen := 1, 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
		} else {
			prevInc = inc
			inc = 1
		}

		currentMax := max(inc>>1, min(prevInc, inc))
		maxLen = max(maxLen, currentMax)

		if maxLen >= k {
			return true
		}
	}
	return false
}
