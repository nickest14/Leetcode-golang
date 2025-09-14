// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/
// Level: Easy

package leetcode

func maxFreqSum(s string) int {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	freq := make(map[byte]int)
	for i := range s {
		freq[s[i]]++
	}

	maxVowel := 0
	maxConsonant := 0

	for ch, count := range freq {
		if vowels[ch] {
			if count > maxVowel {
				maxVowel = count
			}
		} else {
			if count > maxConsonant {
				maxConsonant = count
			}
		}
	}

	return maxVowel + maxConsonant
}
