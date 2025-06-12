// https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/
// Level: Hard

package leetcode

func findKthNumber(n int, k int) int {
	countPrefix := func(p, n int) int {
		count := 0
		cur := p
		nextP := p + 1

		for cur <= n {
			count += min(nextP, n+1) - cur
			cur *= 10
			nextP *= 10
		}
		return count
	}

	cur := 1
	k--

	for k > 0 {
		count := countPrefix(cur, n)
		if k >= count {
			k -= count
			cur++
		} else {
			cur *= 10
			k--
		}
	}

	return cur
}
