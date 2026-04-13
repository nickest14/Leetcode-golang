// https://leetcode.com/problems/minimum-distance-to-the-target-element/
// Level: Easy

package leetcode

func getMinDistance(nums []int, target int, start int) int {
	left := start
	right := start

	for left >= 0 || right < len(nums) {
		if left >= 0 && nums[left] == target {
			return start - left
		}
		if right < len(nums) && nums[right] == target {
			return right - start
		}
		left--
		right++
	}

	return -1
}
