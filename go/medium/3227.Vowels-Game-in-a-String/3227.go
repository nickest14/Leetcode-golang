// https://leetcode.com/problems/vowels-game-in-a-string/
// Level: Medium

package leetcode

func doesAliceWin(s string) bool {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	even := 1
	odd := 0
	parity := 0

	for _, ch := range s {
		if vowels[ch] {
			parity ^= 1
		}

		if parity == 0 {
			even++
		} else {
			odd++
		}
	}

	return even*odd > 0
}
