// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-ii/
// Level: Medium

package leetcode

import "math"

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	if n == 0 {
		return eventTime
	}

	gaps := make([]int, n+1)
	gaps[0] = startTime[0]
	for i := 1; i < n; i++ {
		gaps[i] = startTime[i] - endTime[i-1]
	}
	gaps[n] = eventTime - endTime[n-1]

	largestRight := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		largestRight[i] = int(math.Max(float64(largestRight[i+1]), float64(gaps[i+1])))
	}

	ans := 0
	largestLeft := 0

	for i := 1; i <= n; i++ {
		duration := endTime[i-1] - startTime[i-1]
		canFitLeft := largestLeft >= duration
		canFitRight := largestRight[i] >= duration

		if canFitLeft || canFitRight {
			merged := gaps[i-1] + gaps[i] + duration
			ans = int(math.Max(float64(ans), float64(merged)))
		}

		current := gaps[i-1] + gaps[i]
		ans = int(math.Max(float64(ans), float64(current)))
		largestLeft = int(math.Max(float64(largestLeft), float64(gaps[i-1])))
	}

	return ans
}
