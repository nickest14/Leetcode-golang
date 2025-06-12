// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
// Level: Easy

package leetcode

func maxDifference(s string) int {
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	max_odd := 0
	min_even := 1<<31 - 1
	for _, count := range freq {
		if count%2 == 0 {
			if min_even > count {
				min_even = count
			}
		} else {
			if max_odd < count {
				max_odd = count
			}
		}
	}
	return max_odd - min_even
}
