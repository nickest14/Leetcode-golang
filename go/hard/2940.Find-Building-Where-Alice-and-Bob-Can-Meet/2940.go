// https://leetcode.com/problems/find-building-where-alice-and-bob-can-meet/
// Level: Hard

package leetcode

import (
	"container/heap"
)

type Query struct {
	height int
	index  int
}

type MinHeap []Query

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].height < h[j].height
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Query))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	q := len(queries)
	ans := make([]int, q)
	for i := range ans {
		ans[i] = -1
	}
	deferred := make([][]Query, n)
	pq := &MinHeap{}
	heap.Init((pq))

	for i, q := range queries {
		alice, bob := q[0], q[1]
		if alice > bob {
			alice, bob = bob, alice
		}

		if alice == bob || heights[alice] < heights[bob] {
			ans[i] = bob
		} else {
			deferred[bob] = append(deferred[bob], Query{height: heights[alice], index: i})
		}
	}

	for i := range n {
		for _, q := range deferred[i] {
			heap.Push(pq, q)
		}
		for pq.Len() > 0 && (*pq)[0].height < heights[i] {
			top := heap.Pop(pq).(Query)
			ans[top.index] = i
		}
	}

	return ans
}
