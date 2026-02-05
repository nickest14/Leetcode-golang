// https://leetcode.com/problems/transformed-array/
// Level: Easy

package leetcode

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, num := range nums {
		switch {
		case num > 0:
			ans[i] = nums[(i+num)%n]
		case num < 0:
			ans[i] = nums[((i+num)%n+n)%n]
		default:
			ans[i] = num
		}
	}
	return ans
}
