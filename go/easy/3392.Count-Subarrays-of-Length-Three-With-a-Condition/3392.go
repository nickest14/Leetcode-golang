// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
// Level: Easy

package leetcode

func countSubarrays(nums []int) int {
	ans := 0
	for i := 1; i < len(nums)-1; i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			ans++
		}
	}
	return ans
}
