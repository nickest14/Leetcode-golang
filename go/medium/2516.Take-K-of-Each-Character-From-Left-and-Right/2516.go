// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/
// Level: Medium

package leetcode

func takeCharacters(s string, k int) int {
	count := make([]int, 3)
	for _, c := range s {
		count[c-'a']++
	}
	if min(count[0], count[1], count[2]) < k {
		return -1
	}

	ans := len(s)
	l := 0
	for r := 0; r < len(s); r++ {
		count[s[r]-'a']--

		for min(count[0], count[1], count[2]) < k {
			count[s[l]-'a']++
			l++
		}
		ans = min(ans, len(s)-(r-l+1))
	}

	return ans
}
