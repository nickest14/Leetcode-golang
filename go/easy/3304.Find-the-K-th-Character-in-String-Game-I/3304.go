// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/
// Level: Easy

package leetcode

func kthCharacter(k int) byte {
	word := []byte{'a'}
	for len(word) < k {
		size := len(word)
		for i := 0; i < size; i++ {
			next_char := byte('a' + ((int(word[i]-'a') + 1) % 26))
			word = append(word, next_char)
		}
	}
	return word[k-1]
}
