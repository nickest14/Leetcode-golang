// https://leetcode.com/problems/continuous-subarrays/
// Level: Medium

package leetcode

import "math"

func maxKey(m map[int]int) int {
	max := math.MinInt
	for key := range m {
		if key > max {
			max = key
		}
	}
	return max
}

func minKey(m map[int]int) int {
	min := math.MaxInt
	for key := range m {
		if key < min {
			min = key
		}
	}
	return min
}

func continuousSubarrays(nums []int) int64 {
	ans := 0
	left := 0
	m := make(map[int]int)

	for right := range len(nums) {
		m[nums[right]]++
		for (maxKey(m) - minKey(m)) > 2 {
			m[nums[left]]--
			if m[nums[left]] == 0 {
				delete(m, nums[left])
			}
			left++
		}

		ans += right - left + 1
	}

	return int64(ans)
}
