// https://leetcode.com/problems/three-consecutive-odds/
// Level: Easy

package leetcode

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for _, num := range arr {
		if num%2 == 0 {
			count = 0
		} else {
			count++
			if count == 3 {
				return true
			}
		}
	}
	return false
}
