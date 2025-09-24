// https://leetcode.com/problems/compare-version-numbers/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

func compareVersion(version1, version2 string) int {
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	maxLen := len(v1Parts)
	if len(v2Parts) > maxLen {
		maxLen = len(v2Parts)
	}

	getVersionNumber := func(parts []string, index int) int {
		if index >= len(parts) {
			return 0
		}
		num, _ := strconv.Atoi(parts[index])
		return num
	}

	for i := 0; i < maxLen; i++ {
		v1Num := getVersionNumber(v1Parts, i)
		v2Num := getVersionNumber(v2Parts, i)

		if v1Num > v2Num {
			return 1
		}
		if v1Num < v2Num {
			return -1
		}
	}
	return 0
}
