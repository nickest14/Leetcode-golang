// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/
// Level: Easy

package leetcode

func minimumOperations(nums []int) int {
	countMap := make([]int, 101)

	for i := len(nums) - 1; i >= 0; i-- {
		countMap[nums[i]]++
		if countMap[nums[i]] == 2 {
			return (i + 3) / 3
		}
	}

	return 0
}
