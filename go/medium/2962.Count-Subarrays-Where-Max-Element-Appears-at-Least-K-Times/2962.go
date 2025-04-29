// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/
// Level: Medium

package leetcode

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)
	ans := 0

	maxValue := 0
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}
	left := 0
	maxCount := 0

	for right, num := range nums {
		if num == maxValue {
			maxCount++
		}

		for maxCount >= k {
			ans += n - right
			if nums[left] == maxValue {
				maxCount--
			}
			left++
		}
	}
	return int64(ans)
}
