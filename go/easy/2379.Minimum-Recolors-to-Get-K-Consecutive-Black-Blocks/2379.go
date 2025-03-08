// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// Level: Easy

package leetcode

func minimumRecolors(blocks string, k int) int {
	ans := k
	blackCount := 0
	for i := 0; i < len(blocks); i++ {
		if i-k >= 0 && blocks[i-k] == 'B' {
			blackCount--
		}
		if blocks[i] == 'B' {
			blackCount++
		}
		ans = min(ans, k-blackCount)
	}
	return ans
}
