// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
// Level: Hard

package leetcode

func numWays(words []string, target string) int {
	const mod int = 1e9 + 7

	m := len(words)
	n := len(words[0])
	length := len(target)

	dic := make([]map[byte]int, n)
	for i := 0; i < n; i++ {
		dic[i] = make(map[byte]int)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := words[i][j]
			dic[j][c]++
		}
	}

	memo := make(map[[2]int]int)

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == length {
			return 1
		}
		if n-i < length-j {
			return 0
		}

		if val, exists := memo[[2]int{i, j}]; exists {
			return val
		}

		// Case 1: Skip the current column
		ans := dfs(i+1, j)

		// Case 2: Use the current column if it has the target character
		if count, exists := dic[i][target[j]]; exists {
			ans = (ans + count*dfs(i+1, j+1)) % mod
		}

		memo[[2]int{i, j}] = ans
		return ans
	}

	return dfs(0, 0)
}
