// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
// Level: Medium

package leetcode

func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)
	ans := -1 << 31
	for i := n - 1; i >= 0; i-- {
		next := 0
		if i+k < n {
			next = dp[i+k]
		}
		dp[i] = energy[i] + next
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
