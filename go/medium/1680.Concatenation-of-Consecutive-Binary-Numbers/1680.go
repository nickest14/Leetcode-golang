// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// Level: Medium

package leetcode

import "math/bits"

func concatenatedBinary(n int) int {
	const mod int64 = 1e9 + 7
	ans := int64(0)
	for i := 1; i <= n; i++ {
		shift := bits.Len(uint(i))
		ans = (ans<<shift + int64(i)) % mod
	}
	return int(ans)
}
