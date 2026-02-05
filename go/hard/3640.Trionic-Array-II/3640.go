// https://leetcode.com/problems/trionic-array-ii/
// Level: Hard
// 條件：|nums[i]| <= 10^9，用 sentinel 避免 MinInt64 溢位

package leetcode

func maxSumTrionic(nums []int) int64 {
	const invalid int64 = -1e18
	n := len(nums)
	a, b, c, ans := invalid, invalid, invalid, invalid

	for i := 1; i < n; i++ {
		x0, x1 := int64(nums[i-1]), int64(nums[i])

		na, nb, nc := invalid, invalid, invalid
		if x0 < x1 {
			na = max(x0+x1, a+x1)
			nc = max(c+x1, b+x1)
		} else if x0 > x1 {
			nb = max(b+x1, a+x1)
		}

		ans = max(ans, nc)
		a, b, c = na, nb, nc
	}
	return ans
}
