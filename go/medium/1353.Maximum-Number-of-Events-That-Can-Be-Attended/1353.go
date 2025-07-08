// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/
// Level: Medium

package leetcode

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	minHeap := &IntHeap{}
	heap.Init(minHeap)

	i := 0
	n := len(events)
	ans := 0

	lastDay := 0
	for _, event := range events {
		if event[1] > lastDay {
			lastDay = event[1]
		}
	}

	for day := 1; day <= lastDay; day++ {
		for i < n && events[i][0] == day {
			heap.Push(minHeap, events[i][1])
			i++
		}

		for minHeap.Len() > 0 && (*minHeap)[0] < day {
			heap.Pop(minHeap)
		}

		if minHeap.Len() > 0 {
			heap.Pop(minHeap)
			ans++
		}

		if i == n && minHeap.Len() == 0 {
			break
		}
	}

	return ans
}
