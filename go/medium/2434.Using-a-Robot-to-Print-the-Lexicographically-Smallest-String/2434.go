// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/
// Level: Medium

package leetcode

func robotWithString(s string) string {
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	stack := make([]rune, 0)
	ans := make([]rune, 0)

	minChar := func() rune {
		for i := 'a'; i <= 'z'; i++ {
			if freq[i] > 0 {
				return i
			}
		}
		return 0
	}

	for _, ch := range s {
		stack = append(stack, ch)
		freq[ch]--

		for len(stack) > 0 && stack[len(stack)-1] <= minChar() {
			ans = append(ans, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(ans)
}
