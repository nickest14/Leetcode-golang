// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/
// Level: Medium

package leetcode

func constructDistancedSequence(n int) []int {
	ans := make([]int, 2*n-1)
	used := make([]bool, n+1)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i >= 2*n-1 {
			return true
		}

		for x := n; x > 0; x-- {
			// x > 1, we check two places. Mind index out of bound here.
			// x = 1, we only check one place
			if (x > 1 && (ans[i] != 0 || i+x >= 2*n-1 || ans[i+x] != 0 || used[x])) ||
				(x == 1 && (ans[i] != 0 || used[x])) {
				continue
			}

			if x > 1 {
				ans[i] = x
				ans[i+x] = x
			} else {
				ans[i] = x
			}
			used[x] = true

			nextI := i + 1
			for nextI < 2*n-1 && ans[nextI] != 0 {
				nextI++
			}

			if dfs(nextI) {
				return true
			}

			// backtracking, restore the state
			if x > 1 {
				ans[i] = 0
				ans[i+x] = 0
			} else {
				ans[i] = 0
			}
			used[x] = false
		}

		return false
	}

	dfs(0)

	return ans
}
