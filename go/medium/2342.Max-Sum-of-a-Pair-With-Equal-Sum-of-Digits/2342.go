// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/
// Level: Medium

package leetcode

import "strconv"

func maximumSum(nums []int) int {
	ans := -1
	freq := make(map[int]int)
	digitSum := func(num int) int {
		sum := 0
		strNum := strconv.Itoa(num)
		for _, c := range strNum {
			sum += int(c - '0')
		}
		return sum
	}

	for _, num := range nums {
		ds := digitSum(num)
		if val, exists := freq[ds]; exists {
			ans = max(ans, val+num)
			freq[ds] = max(val, num)
		} else {
			freq[ds] = num
		}
	}
	return ans
}
