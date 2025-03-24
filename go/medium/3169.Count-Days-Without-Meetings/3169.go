// https://leetcode.com/problems/count-days-without-meetings/
// Level: Medium

package leetcode

import "sort"

func countDays(days int, meetings [][]int) int {
	if len(meetings) == 0 {
		return days
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	mergedMeetings := [][]int{}
	curStart, curEnd := meetings[0][0], meetings[0][1]
	for _, meeting := range meetings {
		start, end := meeting[0], meeting[1]
		if start <= curEnd {
			curEnd = max(end, curEnd)
		} else {
			mergedMeetings = append(mergedMeetings, []int{curStart, curEnd})
			curStart, curEnd = start, end
		}
	}
	mergedMeetings = append(mergedMeetings, []int{curStart, curEnd})

	meetingDays := 0
	for _, meeting := range mergedMeetings {
		meetingDays += meeting[1] - meeting[0] + 1
	}
	return days - meetingDays
}
