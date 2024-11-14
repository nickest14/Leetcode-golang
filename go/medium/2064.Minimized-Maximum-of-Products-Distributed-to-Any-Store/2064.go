// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/
// Level: Medium

package leetcode

func minimizedMaximum(n int, quantities []int) int {
	canDistribute := func(x int) bool {
		stores := 0
		for _, q := range quantities {
			stores += (q + x - 1) / x
		}
		return stores <= n
	}

	left, right := 1, 0
	for _, q := range quantities {
		if q > right {
			right = q
		}
	}

	ans := 0
	for left <= right {
		x := (left + right) / 2
		if canDistribute(x) {
			ans = x
			right = x - 1
		} else {
			left = x + 1
		}
	}

	return ans
}
