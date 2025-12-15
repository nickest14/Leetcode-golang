// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/
// Level: Medium

package leetcode

func getDescentPeriods(prices []int) int64 {
	ans := 0
	count := 1
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			count++
		} else {
			ans += count * (count + 1) / 2
			count = 1
		}
	}
	ans += count * (count + 1) / 2
	return int64(ans)
}
