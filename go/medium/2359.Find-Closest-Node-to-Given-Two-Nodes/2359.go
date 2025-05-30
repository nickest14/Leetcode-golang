// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/
// Level: Medium

package leetcode

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)

	distToNode := func(node int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}

		current := node
		distance := 0

		for current != -1 && dist[current] == -1 {
			dist[current] = distance
			distance++
			current = edges[current]
		}

		return dist
	}

	dist1 := distToNode(node1)
	dist2 := distToNode(node2)

	minMax := 1<<31 - 1
	ans := -1
	for i := 0; i < n; i++ {
		if dist1[i] != -1 && dist2[i] != -1 {
			currentMax := max(dist1[i], dist2[i])
			if currentMax < minMax {
				minMax = currentMax
				ans = i
			}
		}
	}
	return ans
}
