// https://leetcode.com/problems/special-array-i/?envType=daily-question&envId=2025-02-01
// Level: Easy

package leetcode

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}
	return true
}
