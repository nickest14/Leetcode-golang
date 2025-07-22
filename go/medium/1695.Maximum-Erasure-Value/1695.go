// https://leetcode.com/problems/maximum-erasure-value/
// Level: Medium

package leetcode

func maximumUniqueSubarray(nums []int) int {
	seen := make(map[int]bool)
	left := 0
	currentSum := 0
	ans := 0

	for right := 0; right < len(nums); right++ {
		for seen[nums[right]] {
			currentSum -= nums[left]
			delete(seen, nums[left])
			left++
		}

		currentSum += nums[right]
		seen[nums[right]] = true
		ans = max(ans, currentSum)
	}

	return ans
}
