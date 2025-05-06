// https://leetcode.com/problems/build-array-from-permutation/
// Level: Easy

package leetcode

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, num := range nums {
		ans[i] = nums[num]
	}
	return ans
}
