// https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/
// Level: Medium

package leetcode

func maximumLength(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	dp := make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			mod := (nums[i] + nums[j]) % k
			currentVal := dp[i][mod]
			if val, exists := dp[j][mod]; exists {
				dp[i][mod] = max(currentVal, val+1)
			} else {
				dp[i][mod] = max(currentVal, 2)
			}

			ans = max(ans, dp[i][mod])
		}
	}

	return ans
}
