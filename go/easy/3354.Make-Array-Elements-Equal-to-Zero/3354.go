// https://leetcode.com/problems/make-array-elements-equal-to-zero/
// Level: Easy

package leetcode

func countValidSelections(nums []int) int {
	simulate := func(i int, direction int) bool {
		arr := make([]int, len(nums))
		copy(arr, nums)

		for i >= 0 && i < len(arr) {
			if arr[i] == 0 {
				i += direction
			} else {
				arr[i]--
				direction *= -1
				i += direction
			}
		}

		for _, val := range arr {
			if val != 0 {
				return false
			}
		}
		return true
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if simulate(i, 1) {
				ans++
			}
			if simulate(i, -1) {
				ans++
			}
		}
	}
	return ans
}
