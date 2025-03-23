// https://leetcode.com/problems/count-the-number-of-complete-components/
// Level: Medium

package leetcode

func countCompleteComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make(map[int]bool)
	ans := 0

	var dfs func(node int, component *map[int]bool)
	dfs = func(node int, component *map[int]bool) {
		(*component)[node] = true
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor, component)
			}
		}
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			component := make(map[int]bool)
			dfs(i, &component)
			isComplete := true
			for node := range component {
				if len(graph[node]) != len(component)-1 {
					isComplete = false
					break
				}
			}
			if isComplete {
				ans++
			}
		}
	}
	return ans
}
