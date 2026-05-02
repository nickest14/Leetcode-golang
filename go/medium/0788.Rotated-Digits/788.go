// https://leetcode.com/problems/rotated-digits/
// Level: Medium

package leetcode

func rotatedDigits(n int) int {
	ans := 0
	for num := 1; num <= n; num++ {
		check := num
		valid := true
		changed := false
		for check > 0 && valid {
			digit := check % 10
			if digit == 3 || digit == 4 || digit == 7 {
				valid = false
			} else if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
				changed = true
			}
			check /= 10
		}
		if valid && changed {
			ans++
		}
	}
	return ans
}
