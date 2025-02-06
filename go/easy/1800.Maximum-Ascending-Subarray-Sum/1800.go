// https://leetcode.com/problems/maximum-ascending-subarray-sum/
// Level: Easy

package leetcode

func maxAscendingSum(nums []int) int {
	ans := 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] >= nums[i] {
			cur = 0
		}
		cur += nums[i]
		ans = max(ans, cur)
	}
	return ans
}
