// https://leetcode.com/problems/check-if-strings-can-be-made-equal-with-operations-ii/
// Level: Medium

package leetcode

func checkStrings(s1 string, s2 string) bool {
	parityCounts := func(s string) (even, odd [26]int) {
		for i := 0; i < len(s); i++ {
			c := s[i] - 'a'
			if i%2 == 0 {
				even[c]++
			} else {
				odd[c]++
			}
		}
		return
	}

	e1, o1 := parityCounts(s1)
	e2, o2 := parityCounts(s2)
	return e1 == e2 && o1 == o2
}
