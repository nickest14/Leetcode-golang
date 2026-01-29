// https://leetcode.com/problems/minimum-cost-path-with-edge-reversals/
// Level: Medium

package leetcode

import (
	"container/heap"
	"math"
)

type Item struct {
	cost int
	node int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func minCost(n int, edges [][]int) int {
	graph := make(map[int][]struct {
		target int
		weight int
	})
	for _, edge := range edges {
		source, target, weight := edge[0], edge[1], edge[2]
		graph[source] = append(graph[source], struct {
			target int
			weight int
		}{target, weight})
		graph[target] = append(graph[target], struct {
			target int
			weight int
		}{source, weight * 2})
	}

	minCosts := make([]int, n)
	for i := range minCosts {
		minCosts[i] = math.MaxInt
	}
	minCosts[0] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{cost: 0, node: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*Item)
		currentCost, currentNode := cur.cost, cur.node

		if currentCost > minCosts[currentNode] {
			continue
		}

		if currentNode == n-1 {
			return currentCost
		}

		for _, neighbor := range graph[currentNode] {
			newCost := currentCost + neighbor.weight
			if newCost < minCosts[neighbor.target] {
				minCosts[neighbor.target] = newCost
				heap.Push(pq, &Item{cost: newCost, node: neighbor.target})
			}
		}
	}

	return -1
}
