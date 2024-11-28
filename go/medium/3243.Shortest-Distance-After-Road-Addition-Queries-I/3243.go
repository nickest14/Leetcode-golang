// https://leetcode.com/problems/shortest-distance-after-road-addition-queries-i/
// Level: Medium

package leetcode

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	distances := make([]int, n)
	for i := range distances {
		distances[i] = n - 1 - i
	}
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i+1] = append(graph[i+1], i)
	}
	var updateDistances func(current int)
	updateDistances = func(current int) {
		newDist := distances[current] + 1
		for _, neighbor := range graph[current] {
			if distances[neighbor] < newDist {
				continue
			}
			distances[neighbor] = newDist
			updateDistances(neighbor)
		}
	}

	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		source, target := query[0], query[1]
		graph[target] = append(graph[target], source)
		if distances[target]+1 < distances[source] {
			distances[source] = distances[target] + 1
		}
		updateDistances(source)
		ans = append(ans, distances[0])
	}
	return ans
}
