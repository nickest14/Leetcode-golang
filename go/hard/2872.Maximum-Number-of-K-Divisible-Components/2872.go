// https://leetcode.com/problems/maximum-number-of-k-divisible-components/
// Level: Hard

package leetcode

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adj := make([][]int, n)
	visited := make([]bool, n)
	visited[0] = true
	ans := 0

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// DFS
	var dfs func(node int) int
	dfs = func(node int) int {
		leaf := values[node]
		visited[node] = true
		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				leaf += dfs(neighbor)
			}
		}
		if leaf%k == 0 {
			ans++
			leaf = 0
		}
		return leaf
	}

	dfs(0)
	if ans == 0 {
		return 1
	}
	return ans
}
