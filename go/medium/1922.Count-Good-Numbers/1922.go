// https://leetcode.com/problems/count-good-numbers/
// Level: Medium

package leetcode

func countGoodNumbers(n int64) int {
	MOD := 1000000007
	evenCount := int((n + 1) / 2)
	oddCount := int(n / 2)

	pow := func(base, exp int) int {
		result := 1
		for exp > 0 {
			if exp%2 == 1 {
				result = (result * base) % MOD
			}
			base = (base * base) % MOD
			exp /= 2
		}
		return result
	}

	evenResult := pow(5, evenCount)
	oddResult := pow(4, oddCount)
	return int((evenResult * oddResult) % MOD)
}
