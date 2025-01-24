// https://leetcode.com/problems/find-eventual-safe-states/

// Level: Medium

package leetcode

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	safe := make(map[int]bool, n)
	var dfs func(int) bool
	dfs = func(node int) bool {
		if val, ok := safe[node]; ok {
			return val
		}
		safe[node] = false
		for _, neighbor := range graph[node] {
			if !dfs(neighbor) {
				return false
			}
		}
		safe[node] = true
		return true
	}

	var ans []int
	for i := 0; i < n; i++ {
		if dfs(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
