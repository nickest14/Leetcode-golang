// https://leetcode.com/problems/power-of-four/
// Level: Easy

package leetcode

func isPowerOfFour(n int) bool {
	return n%3 == 1 && n&(n-1) == 0
}
