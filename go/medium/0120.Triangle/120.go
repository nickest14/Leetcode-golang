// https://leetcode.com/problems/triangle/

// Level: Medium

package leetcode

func minimumTotal(triangle [][]int) int {
	level := len(triangle)
	dp := make([]int, len(triangle[level-1]))
	copy(dp, triangle[level-1])

	for layer := level - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			dp[i] = min(dp[i], dp[i+1]) + triangle[layer][i]
		}
	}

	return dp[0]
}
