// https://leetcode.com/problems/count-number-of-maximum-bitwise-or-subsets/
// Level: Medium

package leetcode

func countMaxOrSubsets(nums []int) int {
	maxOr, count := 0, 0
	for _, num := range nums {
		maxOr |= num
	}
	var backTrack func(index, currentOr int)
	backTrack = func(index int, currentOr int) {
		if currentOr == maxOr {
			count++
		}
		for i := index; i < len(nums); i++ {
			newOr := currentOr | nums[i]
			if newOr <= maxOr {
				backTrack(i+1, newOr)
			}
		}
	}

	backTrack(0, 0)

	return count
}
