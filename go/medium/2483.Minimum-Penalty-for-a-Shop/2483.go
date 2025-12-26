//https://leetcode.com/problems/minimum-penalty-for-a-shop/
// Level: Medium

package leetcode

func bestClosingTime(customers string) int {
	n := len(customers)
	prefixN := make([]int, n+1)
	postfixN := make([]int, n+1)

	for i := 1; i <= n; i++ {
		prefixN[i] = prefixN[i-1]
		if customers[i-1] == 'N' {
			prefixN[i]++
		}
	}

	for i := n - 1; i >= 0; i-- {
		postfixN[i] = postfixN[i+1]
		if customers[i] == 'Y' {
			postfixN[i]++
		}
	}

	minPenalty := n + 1
	idx := 0
	for i := 0; i <= n; i++ {
		penalty := prefixN[i] + postfixN[i]
		if penalty < minPenalty {
			minPenalty = penalty
			idx = i
		}
	}

	return idx
}
