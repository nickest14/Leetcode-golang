// https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
// Level: Medium

package leetcode

func maxUniqueSplit(s string) int {
	var fn func(i int, seen map[string]bool) int
	fn = func(i int, seen map[string]bool) int {
		ans := 0
		if i < len(s) {
			for j := i + 1; j <= len(s); j++ {
				sub := s[i:j]
				if !seen[sub] {
					seen[sub] = true
					ans = max(ans, 1+fn(j, seen))
					delete(seen, sub)
				}
			}
		}
		return ans
	}

	return fn(0, make(map[string]bool))
}
