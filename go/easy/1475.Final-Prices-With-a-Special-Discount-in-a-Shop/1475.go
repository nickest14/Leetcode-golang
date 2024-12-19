// https://leetcode.com/problems/final-prices-with-a-special-discount-in-a-shop/
// Level: Easy

package leetcode

func finalPrices(prices []int) []int {
	stack := []int{}

	for i, price := range prices {
		for len(stack) > 0 && prices[stack[len(stack)-1]] >= price {
			top := stack[len(stack)-1]
			prices[top] -= price
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return prices
}
