// https://leetcode.com/problems/1-bit-and-2-bit-characters/
// Level: Easy

package leetcode

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1 {
		i += bits[i] + 1
	}
	return i == n-1
}
