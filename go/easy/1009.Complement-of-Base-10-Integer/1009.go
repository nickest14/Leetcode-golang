// https://leetcode.com/problems/complement-of-base-10-integer/
// Level: Easy

package leetcode

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	mask := n
	mask |= mask >> 1
	mask |= mask >> 2
	mask |= mask >> 4
	mask |= mask >> 8
	mask |= mask >> 16
	return mask ^ n
}
