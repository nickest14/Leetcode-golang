// https://leetcode.com/problems/bitwise-xor-of-all-pairings/
// Level: Medium

package leetcode

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums1)%2 == 1 {
		for _, num := range nums2 {
			ans ^= num
		}
	}
	if len(nums2)%2 == 1 {
		for _, num := range nums1 {
			ans ^= num
		}
	}
	return ans
}
