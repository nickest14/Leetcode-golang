// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/
// Level: Easy

package leetcode

func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1
}
