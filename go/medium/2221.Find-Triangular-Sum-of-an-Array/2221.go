// https://leetcode.com/problems/find-triangular-sum-of-an-array/
// Level: Medium

package leetcode

func triangularSum(nums []int) int {
	n := len(nums)

	for size := n; size > 1; size-- {
		for i := 0; i < size-1; i++ {
			nums[i] = (nums[i] + nums[i+1]) % 10
		}
	}
	return nums[0]
}
