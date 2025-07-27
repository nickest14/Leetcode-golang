// https://leetcode.com/problems/count-hills-and-valleys-in-an-array/
// Level: Easy

package leetcode

func countHillValley(nums []int) int {
	ans := 0
	left := 0

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			if (nums[left] < nums[i] && nums[i] > nums[i+1]) ||
				(nums[left] > nums[i] && nums[i] < nums[i+1]) {
				ans++
			}
			left = i
		}
	}
	return ans
}
