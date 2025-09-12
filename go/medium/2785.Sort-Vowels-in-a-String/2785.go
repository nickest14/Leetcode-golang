// https://leetcode.com/problems/sort-vowels-in-a-string/
// Level: Medium

package leetcode

import (
	"sort"
	"strings"
)

func sortVowels(s string) string {
	vowels := "aeiouAEIOU"
	vowelList := []byte{}
	for i := range s {
		if strings.Contains(vowels, string(s[i])) {
			vowelList = append(vowelList, s[i])
		}
	}
	sort.Slice(vowelList, func(i, j int) bool {
		return vowelList[i] < vowelList[j]
	})

	ans := []byte(s)
	vowelIdx := 0
	for i := range ans {
		if strings.Contains(vowels, string(ans[i])) {
			ans[i] = vowelList[vowelIdx]
			vowelIdx++
		}
	}
	return string(ans)
}
