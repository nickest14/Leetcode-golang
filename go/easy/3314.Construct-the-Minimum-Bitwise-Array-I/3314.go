// https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/solutions/7508614/100-easy-problem-with-easy-approach-with-yk0g/
// Level: Easy

package leetcode

func minBitwiseArray(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		if num&1 != 0 {
			bit := ((num + 1) & ^num) >> 1
			ans = append(ans, num&^bit)
		} else {
			ans = append(ans, -1)
		}
	}
	return ans
}
