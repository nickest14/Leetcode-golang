// https://leetcode.com/problems/greatest-sum-divisible-by-three/
// Level: Medium

package leetcode

import "math"

func maxSumDivThree(nums []int) int {
	dp := []int{0, math.MinInt, math.MinInt}

	for _, num := range nums {
		tmp := make([]int, 3)
		copy(tmp, dp)

		for r := 0; r < 3; r++ {
			newR := (r + num) % 3
			if dp[r]+num > tmp[newR] {
				tmp[newR] = dp[r] + num
			}
		}
		dp = tmp
	}
	return dp[0]
}
