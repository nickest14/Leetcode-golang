// https://leetcode.com/problems/trionic-array-i/
// Level: Easy

package leetcode

func isTrionic(nums []int) bool {
	isDecreasing := func(a int, b int) bool {
		if a == 0 || b == len(nums)-1 {
			return false
		}
		for i := a; i < b; i++ {
			if nums[i] <= nums[i+1] {
				return false
			}
		}
		return true
	}

	n := len(nums)
	peak := n - 1
	valley := 0

	for i := 0; i < n-1; i++ {
		if peak == n-1 && nums[i] >= nums[i+1] {
			peak = i
		}
		if valley == 0 && nums[n-1-i] <= nums[n-2-i] {
			valley = n - 1 - i
		}
		if peak < valley {
			return isDecreasing(peak, valley)
		}
	}

	return false
}
