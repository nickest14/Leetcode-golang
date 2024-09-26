// https://leetcode.com/problems/find-the-length-of-the-longest-common-prefix/
// Level: Medium

package leetcode

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	dict1 := make(map[string]int)
	maxLen := 0
	for _, num := range arr1 {
		str1 := ""
		for _, ch := range strconv.Itoa(num) {
			str1 += string(ch)
			dict1[str1]++
		}
	}

	for _, num := range arr2 {
		str2 := ""
		for _, ch := range strconv.Itoa(num) {
			str2 += string(ch)
			if _, exists := dict1[str2]; exists {
				if len(str2) > maxLen {
					maxLen = len(str2)
				}
			}
		}
	}

	return maxLen
}
