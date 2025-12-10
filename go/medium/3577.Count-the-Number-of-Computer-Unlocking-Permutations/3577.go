// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/
// Level: Medium

package leetcode

func countPermutations(complexity []int) int {
	const MOD = 1000000007
	n := len(complexity)
	first := complexity[0]

	for i := 1; i < n; i++ {
		if complexity[i] <= first {
			return 0
		}
	}

	ans := 1
	for i := 2; i < n; i++ {
		ans = (ans * i) % MOD
	}

	return ans
}
