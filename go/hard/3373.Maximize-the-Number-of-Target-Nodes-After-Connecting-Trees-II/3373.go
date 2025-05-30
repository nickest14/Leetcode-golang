// https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/

// Level: Hard

package leetcode

func maxTargetNodes(edges1 [][]int, edges2 [][]int) []int {
	bfs := func(graph [][]int) ([]int, []int) {
		n := len(graph)
		colorCount := []int{0, 0}
		nodeColor := make([]int, n)
		visited := make([]bool, n)

		queue := [][2]int{{0, 0}}
		visited[0] = true

		for len(queue) > 0 {
			node := queue[0][0]
			color := queue[0][1]
			queue = queue[1:]

			nodeColor[node] = color
			colorCount[color]++
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, [2]int{neighbor, 1 - color})
				}
			}
		}

		return colorCount, nodeColor
	}

	n1 := len(edges1) + 1
	n2 := len(edges2) + 1
	adj1 := make([][]int, n1)
	adj2 := make([][]int, n2)

	for i := range adj1 {
		adj1[i] = make([]int, 0)
	}
	for i := range adj2 {
		adj2[i] = make([]int, 0)
	}

	for _, edge := range edges1 {
		u, v := edge[0], edge[1]
		adj1[u] = append(adj1[u], v)
		adj1[v] = append(adj1[v], u)
	}
	for _, edge := range edges2 {
		u, v := edge[0], edge[1]
		adj2[u] = append(adj2[u], v)
		adj2[v] = append(adj2[v], u)
	}

	color1, nodeColor1 := bfs(adj1)
	color2, _ := bfs(adj2)

	maxColor2 := 0
	for _, count := range color2 {
		if count > maxColor2 {
			maxColor2 = count
		}
	}

	ans := make([]int, n1)
	for i := 0; i < n1; i++ {
		ans[i] = color1[nodeColor1[i]] + maxColor2
	}

	return ans
}
