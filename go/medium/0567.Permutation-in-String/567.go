// https://leetcode.com/problems/permutation-in-string/
// Level: Medium

package leetcode

func equalMaps(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	countMap := func(s string) map[rune]int {
		m := make(map[rune]int)
		for _, c := range s {
			m[c]++
		}
		return m
	}

	s1Count := countMap(s1)
	windowCount := make(map[rune]int)
	for ind, c := range s2 {
		windowCount[c]++
		if ind >= len(s1) {
			leftChar := rune(s2[ind-len(s1)])
			if windowCount[leftChar] == 1 {
				delete(windowCount, leftChar)
			} else {
				windowCount[leftChar]--
			}
		}
		if equalMaps(s1Count, windowCount) {
			return true
		}
	}

	return false
}
