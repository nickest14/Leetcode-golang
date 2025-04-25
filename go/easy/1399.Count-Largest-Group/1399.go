// https://leetcode.com/problems/count-largest-group/
// Level: Easy

package leetcode

import "strconv"

func countLargestGroup(n int) int {
	prevSum := 0
	groups := make(map[int]int)
	for i := 1; i <= n; i++ {
		if i%10 == 0 {
			s := strconv.Itoa(i)
			prevSum = 0
			for _, c := range s {
				prevSum += int(c - '0')
			}
		} else {
			prevSum += 1
		}
		groups[prevSum]++
	}
	maxGroup := 0
	for _, v := range groups {
		if v > maxGroup {
			maxGroup = v
		}
	}

	ans := 0
	for _, v := range groups {
		if v == maxGroup {
			ans++
		}
	}
	return ans
}
