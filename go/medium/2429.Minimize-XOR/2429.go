// https://leetcode.com/problems/minimize-xor/
// Level: Medium

package leetcode

import "math/bits"

func minimizeXor(num1 int, num2 int) int {
	count1 := bits.OnesCount(uint(num1))
	count2 := bits.OnesCount(uint(num2))

	ans := num1
	for i := 0; i < 32; i++ {
		if count1 > count2 && (num1&(1<<i)) != 0 {
			ans ^= (1 << i)
			count1--
		} else if count1 < count2 && (num1&(1<<i)) == 0 {
			ans ^= (1 << i)
			count1++
		} else if count1 == count2 {
			break
		}
	}
	return ans
}
