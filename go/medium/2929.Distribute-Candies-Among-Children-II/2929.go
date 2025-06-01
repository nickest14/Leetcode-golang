// https://leetcode.com/problems/distribute-candies-among-children-ii/
// Level: Medium

package leetcode

func distributeCandies(n int, limit int) int64 {
	// Computes C(x,2)
	c2 := func(x int) int {
		if x >= 2 {
			return x * (x - 1) / 2
		}
		return 0
	}

	total := (n + 2) * (n + 1) / 2 // C(n+2, 2)

	x1 := n - limit + 1 // Count with a > limit
	t1 := c2(x1)
	x2 := n - 2*limit // Count with a & b > limit
	t2 := c2(x2)
	x3 := n - 3*limit - 1 // Count with a & b & c > limit
	t3 := c2(x3)

	return int64(total - 3*t1 + 3*t2 - t3)
}
