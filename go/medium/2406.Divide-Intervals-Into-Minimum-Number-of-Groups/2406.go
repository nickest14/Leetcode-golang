// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/
// Level: Medium

package leetcode

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x

}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := &MinHeap{}
	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if ans.Len() > 0 && (*ans)[0] < start {
			heap.Pop(ans)
		}
		heap.Push(ans, end)
	}
	return ans.Len()
}
