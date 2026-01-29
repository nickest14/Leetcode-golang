// https://leetcode.com/problems/minimum-cost-to-convert-string-i/
// Level: Medium

package leetcode

import (
	"container/heap"
	"math"
)

type Item struct {
	dist int
	node int
}

type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].dist < h[j].dist
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	graph := make([][][2]int, 26)
	for i := 0; i < len(original); i++ {
		x := int(original[i] - 'a')
		y := int(changed[i] - 'a')
		graph[x] = append(graph[x], [2]int{y, cost[i]})
	}

	dijkstra := func(start int) []int {
		dist := make([]int, 26)
		for i := range dist {
			dist[i] = math.MaxInt
		}
		dist[start] = 0
		pq := &MinHeap{}
		heap.Init(pq)
		heap.Push(pq, &Item{dist: 0, node: start})
		for pq.Len() > 0 {
			cur := heap.Pop(pq).(*Item)
			d, u := cur.dist, cur.node
			if d > dist[u] {
				continue
			}
			for _, e := range graph[u] {
				v, w := e[0], e[1]
				nd := d + w
				if nd < dist[v] {
					dist[v] = nd
					heap.Push(pq, &Item{dist: nd, node: v})
				}
			}
		}
		return dist
	}

	compute := make([][]int, 26)
	for i := 0; i < 26; i++ {
		compute[i] = dijkstra(i)
	}

	var ans int64
	for i := 0; i < len(source); i++ {
		si := int(source[i] - 'a')
		ti := int(target[i] - 'a')
		d := compute[si][ti]
		if d == math.MaxInt {
			return -1
		}
		ans += int64(d)
	}
	return ans
}
