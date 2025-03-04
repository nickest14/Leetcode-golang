// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
// Level: Medium

package leetcode

func checkPowersOfThree(n int) bool {
	power := 1
	for n > power {
		power *= 3
	}

	for n > 0 && power > 0 {
		if n >= power {
			n -= power
		}
		power /= 3
	}

	return n == 0
}
