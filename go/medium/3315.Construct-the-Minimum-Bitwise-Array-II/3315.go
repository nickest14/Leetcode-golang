// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/
// Level: Medium

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
