// https://leetcode.com/problems/reschedule-meetings-for-maximum-free-time-i/
// Level: Medium

package leetcode

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	count := len(startTime)
	ans := 0
	prefixSum := make([]int, count+1)

	for i := 0; i < count; i++ {
		prefixSum[i+1] = prefixSum[i] + (endTime[i] - startTime[i])
	}

	for i := k - 1; i < count; i++ {
		occupied := prefixSum[i+1] - prefixSum[i+1-k]
		windowEnd := eventTime
		if i != count-1 {
			windowEnd = startTime[i+1]
		}
		windowStart := 0
		if i != k-1 {
			windowStart = endTime[i-k]
		}
		freeTime := windowEnd - windowStart - occupied
		if freeTime > ans {
			ans = freeTime
		}
	}

	return ans
}
