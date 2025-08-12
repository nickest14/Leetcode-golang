// https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/
// Level: Medium

package leetcode

func numberOfWays(n int, x int) int {
	mod := 1000000007
	powers := []int{}
	for i := 1; ; i++ {
		p := 1
		for k := 0; k < x; k++ {
			p *= i
		}
		if p > n {
			break
		}
		powers = append(powers, p)
	}

	dp := make([]int, n+1)
	dp[0] = 1
	for _, p := range powers {
		for s := n; s >= p; s-- {
			dp[s] = (dp[s] + dp[s-p]) % mod
		}
	}
	return int(dp[n])
}
