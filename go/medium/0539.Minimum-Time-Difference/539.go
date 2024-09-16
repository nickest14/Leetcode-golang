// https://leetcode.com/problems/minimum-time-difference/
// Level: Medium

package leetcode

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	times := make([]int, len(timePoints))
	for i, time := range timePoints {
		hours, _ := strconv.Atoi(time[0:2])
		minutes, _ := strconv.Atoi(time[3:5])
		times[i] = hours*60 + minutes
	}
	sort.Ints(times)
	ans := 1440 - (times[len(times)-1] - times[0])
	for i := 1; i < len(times); i++ {
		ans = min(ans, times[i]-times[i-1])
	}
	return ans
}
