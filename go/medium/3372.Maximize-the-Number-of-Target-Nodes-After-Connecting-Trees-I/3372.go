// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/

// Level: Medium

package leetcode

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	var dfs func(node, rem, par int, adj [][]int) int
	dfs = func(node, rem, par int, adj [][]int) int {
		cnt := 1
		if rem > 0 {
			for _, neighbor := range adj[node] {
				if neighbor != par {
					cnt += dfs(neighbor, rem-1, node, adj)
				}
			}
		}
		return cnt
	}

	n1 := len(edges1) + 1
	n2 := len(edges2) + 1
	adj1 := make([][]int, n1)
	for i := 0; i < n1; i++ {
		adj1[i] = []int{}
	}
	for _, edge := range edges1 {
		u := edge[0]
		v := edge[1]
		adj1[u] = append(adj1[u], v)
		adj1[v] = append(adj1[v], u)
	}

	adj2 := make([][]int, n2)
	for i := 0; i < n2; i++ {
		adj2[i] = []int{}
	}
	for _, edge := range edges2 {
		u := edge[0]
		v := edge[1]
		adj2[u] = append(adj2[u], v)
		adj2[v] = append(adj2[v], u)
	}

	ans := make([]int, n1)
	if k == 0 {
		for i := 0; i < n1; i++ {
			ans[i] = 1
		}
		return ans
	}

	maxi := 0
	for i := 0; i < n2; i++ {
		cnt := dfs(i, k-1, -1, adj2)
		if cnt > maxi {
			maxi = cnt
		}
	}

	for i := 0; i < n1; i++ {
		cnt := dfs(i, k, -1, adj1)
		ans[i] = cnt + maxi
	}

	return ans
}
