// https://leetcode.com/problems/trapping-rain-water-ii/
// Level: Hard

package leetcode

import (
	"container/heap"
)

type cell struct {
	height, x, y int
}

type minHeap []cell

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].height < h[j].height
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(cell))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	h := &minHeap{}
	heap.Init(h)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{height: heightMap[i][j], x: i, y: j})
				heightMap[i][j] = -1
			}
		}
	}

	level := 0
	ans := 0
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for h.Len() > 0 {
		c := heap.Pop(h).(cell)
		level = max(level, c.height)
		for _, dir := range directions {
			ni, nj := c.x+dir[0], c.y+dir[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && heightMap[ni][nj] != -1 {
				heap.Push(h, cell{height: heightMap[ni][nj], x: ni, y: nj})
				if heightMap[ni][nj] < level {
					ans += level - heightMap[ni][nj]
				}
				heightMap[ni][nj] = -1
			}
		}
	}
	return ans
}
