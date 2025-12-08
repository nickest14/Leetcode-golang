// https://leetcode.com/problems/count-square-sum-triples/
// Level: Easy

package leetcode

import "math"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countTriples(n int) int {
	ans := 0
	maxU := int(math.Sqrt(float64(n)))
	for u := 2; u <= maxU; u++ {
		for v := 1; v < u; v++ {
			if (u-v)%2 == 0 || gcd(u, v) != 1 {
				continue
			}
			c := u*u + v*v
			if c > n {
				continue
			}
			ans += 2 * (n / c)
		}
	}
	return ans
}
