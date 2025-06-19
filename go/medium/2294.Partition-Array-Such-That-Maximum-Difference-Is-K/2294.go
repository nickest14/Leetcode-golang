// https://leetcode.com/problems/partition-array-such-that-maximum-difference-is-k/
// Level: Medium

package leetcode

import "sort"

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	minNum := nums[0]
	ans := 1
	for _, num := range nums[1:] {
		if num-minNum > k {
			ans++
			minNum = num
		}
	}
	return ans
}
