// https://leetcode.com/problems/minimum-deletions-to-make-string-k-special/
// Level: Medium

package leetcode

func minimumDeletions(word string, k int) int {
	ans := len(word)
	freq := make(map[rune]int)
	for _, c := range word {
		freq[c]++
	}

	for _, refCount := range freq {
		changed := 0
		for _, count := range freq {
			if count < refCount {
				changed += count
			} else if count-refCount > k {
				changed += count - refCount - k
			}
		}
		if changed < ans {
			ans = changed
		}
	}
	return ans
}
