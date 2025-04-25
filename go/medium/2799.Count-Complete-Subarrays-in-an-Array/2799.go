// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Level: Medium

package leetcode

func countCompleteSubarrays(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	distinctMap := make(map[int]bool)
	for _, num := range nums {
		distinctMap[num] = true
	}
	distinctCount := len(distinctMap)

	ans := 0
	for i := 0; i < n; i++ {
		seen := make(map[int]bool)
		for j := i; j < n; j++ {
			seen[nums[j]] = true
			if len(seen) == distinctCount {
				ans += n - j
				break
			}
		}
	}
	return ans
}
