// https://leetcode.com/problems/minimum-element-after-replacement-with-digit-sum/
// Level: Easy

package leetcode

func minElement(nums []int) int {
	digitSum := func(n int) int {
		total := 0
		for n > 0 {
			total += n % 10
			n /= 10
		}
		return total
	}

	min := digitSum(nums[0])
	for i := 1; i < len(nums); i++ {
		sum := digitSum(nums[i])
		if sum < min {
			min = sum
		}
	}
	return min
}
