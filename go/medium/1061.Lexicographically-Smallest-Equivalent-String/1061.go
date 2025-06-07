// https://leetcode.com/problems/construct-k-palindrome-strings/
// Level: Medium

package leetcode

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
	for i := range parent {
		parent[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if x != parent[x] {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	for i := 0; i < len(s1); i++ {
		charS1 := int(s1[i] - 'a')
		charS2 := int(s2[i] - 'a')
		parent1 := find(charS1)
		parent2 := find(charS2)

		if parent1 < parent2 {
			parent[parent2] = parent1
		} else {
			parent[parent1] = parent2
		}
	}

	ans := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		index := int(baseStr[i] - 'a')
		ans[i] = byte(find(index) + 'a')
	}

	return string(ans)
}
