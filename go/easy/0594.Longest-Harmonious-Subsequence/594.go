// https://leetcode.com/problems/longest-harmonious-subsequence/
// Level: Easy

package leetcode

func findLHS(nums []int) int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	ans := 0
	for _, num := range nums {
		if count, exists := counter[num+1]; exists {
			currentLength := counter[num] + count
			if currentLength > ans {
				ans = currentLength
			}
		}
	}

	return ans
}
