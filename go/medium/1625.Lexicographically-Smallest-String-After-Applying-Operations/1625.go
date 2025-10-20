// https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/
// Level: Medium

package leetcode

func findLexSmallestString(s string, a int, b int) string {
	visited := map[string]bool{s: true}
	ans := s
	q := []string{s}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur < ans {
			ans = cur
		}

		runes := []rune(cur)
		for i := 1; i < len(runes); i += 2 {
			runes[i] = rune(((int(runes[i]-'0') + a) % 10) + '0')
		}
		added := string(runes)
		if !visited[added] {
			visited[added] = true
			q = append(q, added)
		}

		rotated := cur[len(cur)-b:] + cur[:len(cur)-b]
		if !visited[rotated] {
			visited[rotated] = true
			q = append(q, rotated)
		}
	}
	return ans
}
