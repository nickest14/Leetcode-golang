// https://leetcode.com/problems/vowel-spellchecker/
// Level: Medium

package leetcode

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	maskVowels := func(s string) string {
		b := []rune(s)
		for i, ch := range b {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
				b[i] = '#'
			}
		}
		return string(b)
	}

	exact := make(map[string]struct{}, len(wordlist))
	lowerMap := make(map[string]string, len(wordlist))
	vowelMap := make(map[string]string, len(wordlist))

	for _, w := range wordlist {
		exact[w] = struct{}{}
		wl := strings.ToLower(w)
		if _, ok := lowerMap[wl]; !ok {
			lowerMap[wl] = w
		}
		masked := maskVowels(wl)
		if _, ok := vowelMap[masked]; !ok {
			vowelMap[masked] = w
		}
	}

	ans := make([]string, len(queries))
	for i, q := range queries {
		if _, ok := exact[q]; ok {
			ans[i] = q
			continue
		}
		ql := strings.ToLower(q)
		if v, ok := lowerMap[ql]; ok {
			ans[i] = v
			continue
		}
		qmask := maskVowels(ql)
		if v, ok := vowelMap[qmask]; ok {
			ans[i] = v
		} else {
			ans[i] = ""
		}
	}
	return ans
}
