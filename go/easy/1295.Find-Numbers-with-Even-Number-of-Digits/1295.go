// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/
// Level: Easy

package leetcode

import "strconv"

func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			ans++
		}
	}
	return ans
}
