// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/
// Level: Hard

package leetcode

func minDeletionSize(strs []string) int {
	m := len(strs[0])
	dp := make([]int, m)
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			valid := true
			for _, row := range strs {
				if row[j] > row[i] {
					valid = false
					break
				}
			}
			if valid {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	maxDP := dp[0]
	for i := 1; i < m; i++ {
		if dp[i] > maxDP {
			maxDP = dp[i]
		}
	}
	return m - maxDP
}
