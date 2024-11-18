// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/
// Level: Hard

package leetcode

import "math"

func shortestSubarray(nums []int, k int) int {
	q := [][2]int{{0, 0}}
	ans := math.MaxInt32
	cur := 0
	for i, num := range nums {
		cur += num
		for len(q) > 0 && cur-q[0][1] >= k {
			ans = min(ans, i+1-q[0][0])
			q = q[1:]
		}
		for len(q) > 0 && cur <= q[len(q)-1][1] {
			q = q[:len(q)-1]
		}
		q = append(q, [2]int{i + 1, cur})
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
