// https://leetcode.com/problems/string-compression-iii/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

func compressedString(word string) string {
	var builder strings.Builder
	count := 1
	curC := word[0]
	for i := 1; i < len(word); i++ {
		if word[i] == curC && count < 9 {
			count++
		} else {
			builder.WriteString(strconv.Itoa(count))
			builder.WriteString(string(curC))
			count = 1
			curC = word[i]
		}
	}
	builder.WriteString(strconv.Itoa(count))
	builder.WriteString(string(curC))
	return builder.String()
}
