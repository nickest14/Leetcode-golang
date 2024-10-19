// https://leetcode.com/problems/find-kth-bit-in-nth-binary-string/
// Level: Medium

package leetcode

func findKthBit(n int, k int) byte {
	revInv := func(s string) string {
		runes := []rune(s)
		// invert
		for i := 0; i < len(s); i++ {
			if runes[i] == '0' {
				runes[i] = '1'
			} else {
				runes[i] = '0'
			}
		}
		// reverse
		for i := 0; i < len(runes)/2; i++ {
			runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
		}
		return string(runes)
	}

	s := "0"
	for n > 1 {
		s = s + "1" + revInv(s)
		n--
	}
	return s[k-1]
}
