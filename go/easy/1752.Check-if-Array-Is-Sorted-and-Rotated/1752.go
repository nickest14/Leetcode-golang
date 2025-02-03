// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/
// Level: Easy

package leetcode

func check(nums []int) bool {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] < nums[(i-1+n)%n] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
