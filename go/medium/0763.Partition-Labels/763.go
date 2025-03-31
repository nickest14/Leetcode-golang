// https://leetcode.com/problems/partition-labels/
// Level: Medium

package leetcode

func partitionLabels(s string) []int {
	lastOccurence := make(map[rune]int)
	for i, char := range s {
		lastOccurence[char] = i
	}

	ans := []int{}
	goal, curLen := 0, 0
	for i, char := range s {
		goal = max(goal, lastOccurence[char])
		curLen++
		if goal == i {
			ans = append(ans, curLen)
			curLen = 0
		}
	}

	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
