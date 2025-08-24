// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Level: Medium

package leetcode

func longestSubarray(nums []int) int {
	left, zeros, ans := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}
