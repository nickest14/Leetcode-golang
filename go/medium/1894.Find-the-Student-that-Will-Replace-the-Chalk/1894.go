// https://leetcode.com/problems/find-the-student-that-will-replace-the-chalk/
// Level: Medium

package leetcode

func chalkReplacer(chalk []int, k int) int {
	totalSum := 0
	for _, c := range chalk {
		totalSum += c
	}
	remain := k % totalSum

	for i, c := range chalk {
		if remain < c {
			return i
		}
		remain -= c
	}
	return 0
}
