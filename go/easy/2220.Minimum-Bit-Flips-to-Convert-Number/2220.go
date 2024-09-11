// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Level: Easy

package leetcode

import "math/bits"

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	return bits.OnesCount(uint(xor))
}

// // Another solution
// func minBitFlips(start int, goal int) int {
// 	xor := start ^ goal
// 	ans := 0
// 	for xor > 0 {
// 		ans += xor & 1 // Check the last bit
// 		xor >>= 1      // Shift the bits to the right
// 	}
// 	return ans
// }
