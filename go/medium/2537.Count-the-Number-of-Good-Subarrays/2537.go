// https://leetcode.com/problems/count-the-number-of-good-subarrays/
// Level: Medium

package leetcode

func countGood(nums []int, k int) int64 {
	n := len(nums)
	left := 0
	pairs := 0
	freq := make(map[int]int)
	ans := 0

	for right := 0; right < n; right++ {
		val := nums[right]
		pairs += freq[val]
		freq[val]++

		for pairs >= k {
			ans += n - right
			out := nums[left]
			freq[out]--
			pairs -= freq[out]
			left++
		}
	}
	return int64(ans)
}
