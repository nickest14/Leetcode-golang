// https://leetcode.com/problems/number-complement/
// Level: Easy

package leetcode

func findComplement(num int) int {
	bitLength := 0
	temp := num
	for temp > 0 {
		bitLength++
		temp = temp >> 1
	}
	mask := (1 << bitLength) - 1
	return num ^ mask
}
