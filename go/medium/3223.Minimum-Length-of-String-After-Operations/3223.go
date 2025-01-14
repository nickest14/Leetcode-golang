// https://leetcode.com/problems/minimum-length-of-string-after-operations/
// Level: Medium

package leetcode

func minimumLength(s string) int {
	freq := make([]int, 26)
	for _, v := range s {
		freq[v-'a']++
	}
	var ans int
	for _, v := range freq {
		if v == 0 {
			continue
		}
		if v%2 == 0 {
			ans += 2
		} else {
			ans++
		}
	}
	return ans
}

// func minimumLength(s string) int {
// 	freq := make(map[rune]int)
// 	for _, c := range s {
// 		freq[c]++
// 	}

// 	ans := 0
// 	for _, v := range freq {
// 		if v%2 == 0 {
// 			ans += 2
// 		} else {
// 			ans++
// 		}
// 	}
// 	return ans
// }
