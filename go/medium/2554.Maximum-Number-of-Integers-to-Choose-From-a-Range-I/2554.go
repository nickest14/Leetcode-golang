// https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/
// Level: Medium

package leetcode

func maxCount(banned []int, n int, maxSum int) int {
	ans := 0
	bannedSet := make(map[int]bool)
	for _, val := range banned {
		bannedSet[val] = true
	}

	totalSum := 0
	for i := 1; i <= n; i++ {
		if bannedSet[i] {
			continue
		}
		totalSum += i
		if totalSum > maxSum {
			break
		}
		ans++
	}
	return ans
}
