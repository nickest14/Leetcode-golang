// https://leetcode.com/problems/solving-questions-with-brainpower/

// Level: Medium

package leetcode

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		points := questions[i][0]
		brainPower := questions[i][1]
		nextQ := i + brainPower + 1
		take := points
		if nextQ < n {
			take += dp[nextQ]
		}
		skip := dp[i+1]
		dp[i] = max(take, skip)
	}

	return int64(dp[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
