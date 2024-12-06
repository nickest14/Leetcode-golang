// https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
// Level: Medium

package leetcode

func canMakeSubsequence(str1 string, str2 string) bool {
	str1Len, str2Len := len(str1), len(str2)
	targetChar := str2[0]

	str1Idx, str2Idx := 0, 0
	for str1Idx < str1Len && str2Idx < str2Len {
		str1Char := str1[str1Idx]
		if str1Char == targetChar || str1Char+1 == targetChar || (str1Char == 'z' && targetChar == 'a') {
			str2Idx++
			if str2Idx < str2Len {
				targetChar = str2[str2Idx]
			}
		}
		str1Idx++
	}

	return str2Idx == str2Len
}
