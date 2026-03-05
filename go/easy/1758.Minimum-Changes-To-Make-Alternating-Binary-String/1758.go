// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string/
// Level: Easy

package leetcode

func minOperations(s string) int {
	countChangesIfStartWith := func(first byte) int {
		changes := 0
		for i := 0; i < len(s); i++ {
			if i%2 == 0 && s[i] != first {
				changes++
			} else if i%2 == 1 && s[i] == first {
				changes++
			}
		}
		return changes
	}
	changes0 := countChangesIfStartWith('0')
	changes1 := countChangesIfStartWith('1')
	return min(changes0, changes1)
}
