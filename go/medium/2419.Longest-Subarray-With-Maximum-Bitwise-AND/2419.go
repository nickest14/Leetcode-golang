// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// Level: Medium

package leetcode

func longestSubarray(nums []int) int {
	ans := 0
	curLen := 0
	maxNum := 0
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	for _, num := range nums {
		if num == maxNum {
			curLen++
		} else {
			ans = max(ans, curLen)
			curLen = 0
		}
	}

	return max(ans, curLen)
}
