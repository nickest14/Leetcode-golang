// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/
// Level: Hard

package leetcode

import "container/list"

func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	type Node struct {
		x, y, removed int
	}

	q := list.New()
	q.PushFront(Node{0, 0, 0})
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	for q.Len() > 0 {
		node := q.Remove(q.Front()).(Node)
		x, y, removed := node.x, node.y, node.removed
		if x == m-1 && y == n-1 {
			return removed
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				visited[nx][ny] = true
				if grid[nx][ny] == 1 {
					q.PushBack(Node{nx, ny, removed + 1})
				} else {
					q.PushFront(Node{nx, ny, removed})
				}
			}
		}
	}
	return -1
}
