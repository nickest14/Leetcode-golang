// https://leetcode.com/problems/number-of-ways-to-split-array/
// Level: Medium

package leetcode

func waysToSplitArray(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	accumulate := 0
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		accumulate += nums[i]
		if 2*accumulate >= total {
			ans++
		}
	}
	return ans
}
