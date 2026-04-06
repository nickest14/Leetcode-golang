// https://leetcode.com/problems/decode-the-slanted-ciphertext/
// Level: Medium

package leetcode

import (
	"math"
	"strings"
)

func decodeCiphertext(encodedText string, rows int) string {
	if len(encodedText) == 0 {
		return encodedText
	}

	cols := int(math.Ceil(float64(len(encodedText)) / float64(rows)))
	var ans []byte

	for start := 0; start < cols; start++ {
		r, c := 0, start
		for r < rows && c < cols {
			idx := r*cols + c
			if idx < len(encodedText) {
				ans = append(ans, encodedText[idx])
			}
			r++
			c++
		}
	}

	return strings.TrimRight(string(ans), " ")
}
