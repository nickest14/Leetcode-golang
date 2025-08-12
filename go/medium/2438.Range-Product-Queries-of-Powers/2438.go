// https://leetcode.com/problems/range-product-queries-of-powers/
// Level: Medium

package leetcode

func productQueries(n int, queries [][]int) []int {
	mod := 1000000007
	powers := []int{}
	power := 1

	for n > 0 {
		if n&1 == 1 {
			powers = append(powers, power)
		}
		power *= 2
		n >>= 1
	}

	size := len(powers)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		dp[i][i] = powers[i]
		for j := i + 1; j < size; j++ {
			dp[i][j] = (dp[i][j-1] * powers[j]) % mod
		}
	}

	ans := []int{}
	for _, query := range queries {
		left, right := query[0], query[1]
		ans = append(ans, dp[left][right])
	}

	return ans
}
