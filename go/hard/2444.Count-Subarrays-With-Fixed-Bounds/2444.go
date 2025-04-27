// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/
// Level: Hard

package leetcode

func countSubarrays(nums []int, minK int, maxK int) int64 {
	ans := 0
	prevMin, prevMax, bad := -1, -1, -1

	for i, num := range nums {
		if num == minK {
			prevMin = i
		}
		if num == maxK {
			prevMax = i
		}
		if num < minK || num > maxK {
			bad = i
		}
		if prevMin != -1 && prevMax != -1 {
			ans += max(0, min(prevMin, prevMax)-bad)
		}
	}

	return int64(ans)
}
