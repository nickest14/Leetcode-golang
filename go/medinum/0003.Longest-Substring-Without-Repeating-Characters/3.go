// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Level: Medium

package leetcode

func lengthOfLongestSubstring(s string) int {
	var record [256]int
	if len(s) == 0 {
		return 0
	}
	left, right, ans := 0, -1, 0
	for left < len(s) {
		if right+1 < len(s) && record[s[right+1]] == 0 {
			record[s[right+1]]++
			right++
		} else {
			record[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	// for index, val := range s {
	// 	fmt.Println(index, val)
	// 	fmt.Printf("\n%q", val)
	// }
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
