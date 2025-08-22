// https://leetcode.com/problems/number-of-zero-filled-subarrays/
// Level: Medium

package leetcode

func zeroFilledSubarray(nums []int) int64 {
	zeroCount := 0
	ans := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		} else {
			ans += zeroCount * (zeroCount + 1) / 2
			zeroCount = 0
		}
	}
	ans += zeroCount * (zeroCount + 1) / 2
	return int64(ans)
}
