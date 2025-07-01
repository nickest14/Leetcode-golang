// https://leetcode.com/problems/find-the-original-typed-string-i/
// Level: Easy

package leetcode

func possibleStringCount(word string) int {
	n := len(word)
	ans := n
	for i := 1; i < n; i++ {
		if word[i] != word[i-1] {
			ans--
		}
	}
	return ans
}
