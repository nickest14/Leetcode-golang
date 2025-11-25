// https://leetcode.com/problems/binary-prefix-divisible-by-5/
// Level: Easy

package leetcode

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	val := 0
	for i, bit := range nums {
		val = (val*2 + bit) % 5
		ans[i] = val == 0
	}
	return ans
}
