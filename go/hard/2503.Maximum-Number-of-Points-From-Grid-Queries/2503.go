// https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/
// Level: Hard

package leetcode

import (
	"container/heap"
	"sort"
)

type HeapItem struct {
	val, row, col int
}

type MinHeap []HeapItem

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxPoints(grid [][]int, queries []int) []int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	visited[0][0] = true

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, HeapItem{grid[0][0], 0, 0})

	points := 0
	ans := make([]int, len(queries))

	type Query struct {
		val, idx int
	}
	sortedQueries := make([]Query, len(queries))
	for i, val := range queries {
		sortedQueries[i] = Query{val, i}
	}
	sort.Slice(sortedQueries, func(i, j int) bool {
		return sortedQueries[i].val < sortedQueries[j].val
	})

	for _, query := range sortedQueries {
		for h.Len() > 0 {
			if (*h)[0].val >= query.val {
				break
			}
			top := heap.Pop(h).(HeapItem)
			points++
			for _, dir := range directions {
				nx, ny := top.row+dir[0], top.col+dir[1]
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && !visited[nx][ny] {
					visited[nx][ny] = true
					heap.Push(h, HeapItem{grid[nx][ny], nx, ny})
				}
			}
		}
		ans[query.idx] = points
	}

	return ans
}
