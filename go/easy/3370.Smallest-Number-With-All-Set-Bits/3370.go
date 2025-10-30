// https://leetcode.com/problems/smallest-number-with-all-set-bits/
// Level: Easy

package leetcode

func smallestNumber(n int) int {
	ans := 0
	i := 0
	for ans < n {
		ans += 1 << i
		i++
	}
	return ans
}
