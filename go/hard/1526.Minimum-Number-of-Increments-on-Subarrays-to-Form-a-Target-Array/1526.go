// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
// Level: Hard

package leetcode

func minNumberOperations(target []int) int {
	ans := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			ans += target[i] - target[i-1]
		}
	}
	return ans
}
