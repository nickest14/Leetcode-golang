// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/
// Level: Medium

package leetcode

func maxIncreasingSubarrays(nums []int) int {
	ans := 0
	inc, prevInc := 1, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
		} else {
			prevInc = inc
			inc = 1
		}

		half := inc >> 1
		m := min(inc, prevInc)
		candidate := max(half, m)
		if candidate > ans {
			ans = candidate
		}
	}
	return ans
}
