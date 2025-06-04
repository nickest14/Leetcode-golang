// https://leetcode.com/problems/find-the-lexicographically-largest-string-from-the-box-i/
// Level: Medium

package leetcode

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	m := len(word)
	maxPart := m - numFriends + 1
	ans := []byte{}
	for i := 0; i < len(word); i++ {
		tmp := []byte{}
		for j := i; j < min(i+maxPart, m); j++ {
			tmp = append(tmp, word[j])
		}
		if string(tmp) > string(ans) {
			ans = tmp
		}
	}
	return string(ans)
}

// for i in range(m):
// tmp: str = ""
// for j in range(i, min(i + max_part, m)):
// 	tmp += word[j]
// if tmp > ans:
// 	ans = tmp

// return ans
