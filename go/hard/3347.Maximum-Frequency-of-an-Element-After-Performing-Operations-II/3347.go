// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/
// Level: Hard

package leetcode

import "sort"

func maxFrequency(nums []int, k int, numOperations int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	presum := make(map[int]int)
	for freqKey, count := range freq {
		presum[freqKey-k] += count
		presum[freqKey+k+1] -= count
	}

	keys := make([]int, 0)
	keySet := make(map[int]bool)

	for key := range presum {
		if !keySet[key] {
			keys = append(keys, key)
			keySet[key] = true
		}
	}

	for key := range freq {
		if !keySet[key] {
			keys = append(keys, key)
			keySet[key] = true
		}
	}

	sort.Ints(keys)

	ans := 0
	total := 0

	for _, key := range keys {
		total += presum[key]
		ops := min(total-freq[key], numOperations)
		ans = max(ans, freq[key]+ops)
	}

	return ans
}
