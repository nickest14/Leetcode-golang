// https://leetcode.com/problems/power-of-three/
// Level: Easy

package leetcode

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
