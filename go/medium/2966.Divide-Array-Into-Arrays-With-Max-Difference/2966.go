// https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/
// Level: Medium

package leetcode

import "sort"

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	for i := 0; i < n; i += 3 {
		if i+2 >= n || nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		ans = append(ans, nums[i:i+3])
	}
	return ans
}
