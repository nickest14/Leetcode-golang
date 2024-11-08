// https://leetcode.com/problems/maximum-xor-for-each-query/
// Level: Medium

package leetcode

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	maxXor := (1 << maximumBit) - 1
	xor := nums[0]
	for i := 1; i < n; i++ {
		xor ^= nums[i]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = xor ^ maxXor
		xor ^= nums[n-1-i]
	}
	return ans
}
