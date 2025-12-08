// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
// Level: Easy

package leetcode

func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}
