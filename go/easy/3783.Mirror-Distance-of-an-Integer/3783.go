// https://leetcode.com/problems/mirror-distance-of-an-integer/
// Level: Easy

package leetcode

func mirrorDistance(n int) int {
	reverseDigits := func(x int) int {
		rev := 0
		for x > 0 {
			rev = rev*10 + x%10
			x /= 10
		}
		return rev
	}
	r := reverseDigits(n)
	return max(n, r) - min(n, r)
}
