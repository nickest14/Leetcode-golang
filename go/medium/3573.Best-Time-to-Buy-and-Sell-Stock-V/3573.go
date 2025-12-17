// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/
// Level: Medium

package leetcode

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumProfit(prices []int, k int) int64 {
	profit := make([]int64, k+1)
	buy := make([]int64, k+1)
	sell := make([]int64, k+1)

	// Use a value within JavaScript safe integer range: [-(2^53-1), 2^53-1]
	// -1e15 is safe and large enough to represent "negative infinity"
	const minVal int64 = -1e15

	for i := 0; i <= k; i++ {
		buy[i] = minVal
		sell[i] = minVal
	}

	for _, price := range prices {
		price64 := int64(price)
		prev := profit[0]
		for i := 1; i <= k; i++ {
			p := profit[i]
			profit[i] = maxInt64(maxInt64(profit[i], buy[i]+price64), sell[i]-price64)
			buy[i] = maxInt64(buy[i], prev-price64)
			sell[i] = maxInt64(sell[i], prev+price64)
			prev = p
		}
	}

	ans := profit[0]
	for i := 1; i <= k; i++ {
		if profit[i] > ans {
			ans = profit[i]
		}
	}
	return ans
}
