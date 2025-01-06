// https://leetcode.com/problems/shifting-letters-ii/
// Level: Medium

package leetcode

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	prefix := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		prefix[start] += 2*direction - 1
		prefix[end+1] -= 2*direction - 1
	}

	ans := []rune(s)
	accumulate := 0
	for i := 0; i < len(s); i++ {
		accumulate += prefix[i]
		for accumulate < 0 {
			accumulate += 26
		}
		ans[i] = rune((int(ans[i]-'a')+accumulate)%26 + int('a'))
	}

	return string(ans)
}
