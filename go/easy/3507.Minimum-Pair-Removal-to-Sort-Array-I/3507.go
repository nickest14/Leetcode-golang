// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/
// Level: Easy

package leetcode

func minimumPairRemoval(nums []int) int {
	isIncreasing := func(nums []int) bool {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				return false
			}
		}
		return true
	}

	ans := 0

	for !isIncreasing(nums) {
		minSum := nums[0] + nums[1]
		index := 0

		for i := 1; i < len(nums)-1; i++ {
			sum := nums[i] + nums[i+1]
			if sum < minSum {
				minSum = sum
				index = i
			}
		}

		nums[index] = minSum
		nums = append(nums[:index+1], nums[index+2:]...)
		ans++
	}
	return ans
}
