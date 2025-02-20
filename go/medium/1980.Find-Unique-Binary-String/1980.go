// https://leetcode.com/problems/find-unique-binary-string/

// Level: Medium

package leetcode

import "strings"

func findDifferentBinaryString(nums []string) string {
	result := []string{}
	for i := range nums {
		if nums[i][i] == '0' {
			result = append(result, "1")
		} else {
			result = append(result, "0")

		}
	}

	return strings.Join(result, "")
}
