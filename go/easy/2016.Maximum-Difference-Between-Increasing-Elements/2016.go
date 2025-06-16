// https://leetcode.com/problems/maximum-difference-between-increasing-elements/
// Level: Easy

package leetcode

func maximumDifference(nums []int) int {
	minNum := nums[0]
	ans := -1

	for _, num := range nums[1:] {
		if num > minNum {
			if num-minNum > ans {
				ans = num - minNum
			}
		} else {
			minNum = num
		}
	}

	return ans
}
