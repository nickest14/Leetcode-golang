// https://leetcode.com/problems/finding-3-digit-even-numbers/
// Level: Easy

package leetcode

import "strconv"

func findEvenNumbers(digits []int) []int {
	freq := make([]int, 10)
	for _, digit := range digits {
		freq[digit]++
	}
	ans := []int{}

	for num := 100; num < 1000; num += 2 {
		candidateFreq := make([]int, 10)
		s := strconv.Itoa(num)
		for i := 0; i < len(s); i++ {
			digit := int(s[i] - '0')
			candidateFreq[digit]++
		}
		valid := true
		for i := 0; i < 10; i++ {
			if candidateFreq[i] > freq[i] {
				valid = false
				break
			}
		}
		if valid {
			ans = append(ans, num)
		}
	}

	return ans
}
