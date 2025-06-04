// https://leetcode.com/problems/candy/
// Level: Hard

package leetcode

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	ans := 0
	for _, c := range candies {
		ans += c
	}
	return ans
}
