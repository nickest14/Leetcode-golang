// https://leetcode.com/problems/count-elements-with-maximum-frequency/
// Level: Easy

package leetcode

func maxFrequencyElements(nums []int) int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	maxFreq := 0
	for _, count := range freq {
		if count > maxFreq {
			maxFreq = count
		}
	}
	ans := 0
	for _, count := range freq {
		if count == maxFreq {
			ans += maxFreq
		}
	}
	return ans
}
