// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/
// Level: Hard

package leetcode

func countSubarrays(nums []int, k int64) int64 {
	ans := 0
	left, partialSum := 0, 0
	for right, num := range nums {
		partialSum += num
		for partialSum*(right-left+1) >= int(k) {
			partialSum -= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return int64(ans)
}
