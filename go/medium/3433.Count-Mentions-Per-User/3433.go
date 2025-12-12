// https://leetcode.com/problems/count-mentions-per-user/
// Level: Medium

package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
	ans := make([]int, numberOfUsers)
	onlineTime := make([]int, numberOfUsers)

	sort.Slice(events, func(i, j int) bool {
		timeI, _ := strconv.Atoi(events[i][1])
		timeJ, _ := strconv.Atoi(events[j][1])
		if timeI != timeJ {
			return timeI < timeJ
		}
		return events[i][0] != "MESSAGE" && events[j][0] == "MESSAGE"
	})

	for _, event := range events {
		msg, timeStr, mentions := event[0], event[1], event[2]
		t, _ := strconv.Atoi(timeStr)

		if msg == "MESSAGE" {
			if mentions == "ALL" {
				for i := 0; i < numberOfUsers; i++ {
					ans[i]++
				}
			} else if mentions == "HERE" {
				for i := 0; i < numberOfUsers; i++ {
					if t >= onlineTime[i] {
						ans[i]++
					}
				}
			} else {
				parts := strings.Fields(strings.ReplaceAll(mentions, "id", ""))
				for _, p := range parts {
					idx, _ := strconv.Atoi(p)
					ans[idx]++
				}
			}
		} else {
			idx, _ := strconv.Atoi(mentions)
			onlineTime[idx] = t + 60
		}
	}
	return ans
}
