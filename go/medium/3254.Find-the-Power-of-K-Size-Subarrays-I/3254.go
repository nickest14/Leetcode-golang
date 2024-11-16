// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/
// Level: Medium

package leetcode

func resultsArray(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	n := len(nums)
	ans := make([]int, 0)
	left := 0
	right := 1
	for right < n {
		isConsecutive := nums[right]-nums[right-1] == 1
		if !isConsecutive {
			for left < right && left+k-1 < n {
				ans = append(ans, -1)
				left++
			}
			left = right
		} else if right-left == k-1 {
			ans = append(ans, nums[right])
			left++
		}
		right++
	}

	return ans
}
