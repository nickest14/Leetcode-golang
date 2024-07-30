// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/
// Level: Medium

package leetcode

func minimumDeletions(s string) int {
	var bCount, deletions int
	for _, char := range s {
		if char == 'a' && bCount > 0 {
			bCount--
			deletions++
		} else if char == 'b' {
			bCount++
		}
	}
	return deletions
}
