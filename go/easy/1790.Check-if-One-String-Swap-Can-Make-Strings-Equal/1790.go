// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
// Level: Easy

package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	diff := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	return len(diff) == 2 && s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}
