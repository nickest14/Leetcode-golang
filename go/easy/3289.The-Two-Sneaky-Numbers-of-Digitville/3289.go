// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
// Level: Easy

package leetcode

func getSneakyNumbers(nums []int) []int {
	ans := []int{}
	numToCount := make(map[int]bool)
	for _, num := range nums {
		if numToCount[num] {
			ans = append(ans, num)
			continue
		}
		numToCount[num] = true
	}
	return ans
}
