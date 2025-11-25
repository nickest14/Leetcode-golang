// https://leetcode.com/problems/smallest-integer-divisible-by-k/
// Level: Medium

package leetcode

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	ans := 0
	for i := 1; i <= k; i++ {
		ans = (ans*10 + 1) % k
		if ans == 0 {
			return i
		}
	}
	return -1
}
