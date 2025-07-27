// https://leetcode.com/problems/maximum-score-from-removing-substrings/
// Level: Medium

package leetcode

func maximumGain(s string, x int, y int) int {
	var point, lowerPoint int
	var pattern, lowerPattern string

	if x > y {
		point = x
		lowerPoint = y
		pattern = "ab"
		lowerPattern = "ba"
	} else {
		point = y
		lowerPoint = x
		pattern = "ba"
		lowerPattern = "ab"
	}

	ans := 0
	stack := []rune{}

	for _, c := range s {
		if c == rune(pattern[1]) && len(stack) > 0 && stack[len(stack)-1] == rune(pattern[0]) {
			stack = stack[:len(stack)-1]
			ans += point
		} else {
			stack = append(stack, c)
		}
	}

	lowerStack := []rune{}
	for _, c := range stack {
		if c == rune(lowerPattern[1]) && len(lowerStack) > 0 && lowerStack[len(lowerStack)-1] == rune(lowerPattern[0]) {
			lowerStack = lowerStack[:len(lowerStack)-1]
			ans += lowerPoint
		} else {
			lowerStack = append(lowerStack, c)
		}
	}

	return ans
}
