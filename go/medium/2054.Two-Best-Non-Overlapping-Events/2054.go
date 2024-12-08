// https://leetcode.com/problems/two-best-non-overlapping-events/
// Level: Medium

package leetcode

import "sort"

type EventPoint struct {
	time    int
	isStart bool
	value   int
}

func maxTwoEvents(events [][]int) int {
	eventPoints := []EventPoint{}
	for _, event := range events {
		start, end, value := event[0], event[1], event[2]
		eventPoints = append(eventPoints, EventPoint{start, true, value})
		eventPoints = append(eventPoints, EventPoint{end + 1, false, value})
	}
	sort.Slice(eventPoints, func(i, j int) bool {
		if eventPoints[i].time == eventPoints[j].time {
			return !eventPoints[i].isStart
		}
		return eventPoints[i].time < eventPoints[j].time
	})

	ans := 0
	prevMax := 0
	for _, event := range eventPoints {
		if event.isStart {
			ans = max(ans, prevMax+event.value)
		} else {
			prevMax = max(prevMax, event.value)
		}

	}
	return ans
}
